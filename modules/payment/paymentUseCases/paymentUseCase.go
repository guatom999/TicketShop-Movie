package paymentUseCases

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/payment"
	"github.com/guatom999/TicketShop-Movie/modules/payment/paymentRepositories"
	"github.com/guatom999/TicketShop-Movie/pkg/queue"
	"github.com/guatom999/TicketShop-Movie/utils"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type (
	PaymentUseCaseService interface {
		BuyTicket(pctx context.Context, cfg *config.Config, req *payment.MovieBuyReq) (*payment.BuyticketRes, error)
		CheckOutWithCreditCard(req *payment.CheckOutWithCreditCard) error
		// UploadFileTest(file multipart.File, object string) error
		// BuyTicketConsumer(pctx context.Context, topic string, resCh chan<- *payment.RollBackReserveSeatRes)
	}

	paymentUseCase struct {
		paymentRepo paymentRepositories.PaymentRepositoryService
		cfg         *config.Config
		opnClient   *omise.Client
		bucketName  string
		uploadPath  string
		// cl          *storage.Client
	}

	UploadResponse struct {
		Filename string `json:"filename"`
		URL      string `json:"url"`
	}
)

func NewPaymentUseCase(
	paymentRepo paymentRepositories.PaymentRepositoryService,
	cfg *config.Config,
	opnClient *omise.Client,
	// cli *storage.Client, close client for a while
) PaymentUseCaseService {

	return &paymentUseCase{
		paymentRepo: paymentRepo,
		cfg:         cfg,
		opnClient:   opnClient,
		bucketName:  "ticket-shop-bucket",
		uploadPath:  "ticket-image/",
		// cl:          cli,
	}
}

func (u *paymentUseCase) CheckOutWithCreditCard(req *payment.CheckOutWithCreditCard) error {

	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   req.Price,
		Currency: "thb",
		Card:     req.Token,
	}
	if err := u.opnClient.Do(charge, createCharge); err != nil {
		log.Fatalf("omise clinet failed  %s", err.Error())
		return errors.New("error: omise client failed")
	}

	log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)

	return nil
}

// func PaymentConsumer(pctx context.Context, cfg *config.Config, topic string) *kafka.Conn {
// 	conn := queue.KafkaConn(cfg, topic)

// 	topicConfigs := kafka.TopicConfig{
// 		Topic:             topic,
// 		NumPartitions:     1,
// 		ReplicationFactor: 1,
// 	}

// 	if err := conn.CreateTopics(topicConfigs); err != nil {
// 		log.Printf("Erorr: Create Topic Failed %s", err.Error())
// 		panic(err.Error())
// 	}

// 	return conn

// }

func (u *paymentUseCase) BuyTicketConsumer(pctx context.Context, cfg *config.Config, groupId string, topic string, key string, resCh chan<- *payment.RollBackReserveSeatRes) {
	reader := queue.KafkaReader(u.cfg, topic, groupId)
	defer reader.Close()

	data := new(payment.RollBackReserveSeatRes)

	for {
		select {
		case <-pctx.Done():
			log.Println("BuyTicketConsumer context cancelled")
			close(resCh)
			return
		default:
			msg, err := reader.ReadMessage(pctx)
			if err != nil {
				log.Printf("Error reading message: %s", err.Error())
				close(resCh)
				return
			}

			if string(msg.Key) == key {
				fmt.Println("============================>")
				if err := json.Unmarshal(msg.Value, data); err != nil {
					fmt.Printf("Error: Unmarshal error %s", err.Error())
				}

				resCh <- data
				close(resCh)
				return
			}
		}
	}
}

func (u *paymentUseCase) BuyTicket(pctx context.Context, cfg *config.Config, req *payment.MovieBuyReq) (*payment.BuyticketRes, error) {

	_, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	stage1 := new(payment.RollBackReserveSeatRes)

	resCh := make(chan *payment.RollBackReserveSeatRes)

	go u.BuyTicketConsumer(pctx, cfg, "reserve-seat-res-group", "reserve-seat-res", "payment", resCh)

	req.CustomerId = strings.TrimPrefix(req.CustomerId, "customer:")

	fmt.Println("req.PosterImage is", req.PosterImage)

	if err := u.paymentRepo.ReserveSeat(pctx, cfg, &payment.ReserveSeatReq{
		MovieName: req.MovieName,
		MovieId:   req.MovieId,
		SeatNo:    req.SeatNo,
	}); err != nil {
		return nil, err
	}

	select {
	case res := <-resCh:
		if res != nil {
			stage1 = &payment.RollBackReserveSeatRes{
				MovieId:     res.MovieId,
				Seat_Number: res.Seat_Number,
				Error:       res.Error,
			}
		}
	case <-time.After(time.Second * 10):
		u.paymentRepo.RollBackReserveSeat(pctx, cfg, &payment.RollBackReservedSeatReq{
			MovieId: req.MovieId,
			SeatNo:  req.SeatNo,
		})
		fmt.Println("Timeout waiting for rollback response")
		return nil, errors.New("timeout waiting for rollback response")
	}

	if stage1.Error != "" {
		fmt.Println("stage1.Error", stage1.Error)
		u.paymentRepo.RollBackReserveSeat(pctx, cfg, &payment.RollBackReservedSeatReq{
			MovieId: req.MovieId,
			SeatNo:  req.SeatNo,
		})

		return nil, errors.New("error: failed to reserve seat")
	}

	if err := u.CheckOutWithCreditCard(&payment.CheckOutWithCreditCard{Token: req.Token, Price: req.Price}); err != nil {
		u.paymentRepo.RollBackReserveSeat(pctx, cfg, &payment.RollBackReservedSeatReq{
			MovieId: req.MovieId,
			SeatNo:  req.SeatNo,
		})
		return nil, err
	}

	// var png []byte
	// var fileUrl string
	// reqQrCode := utils.GenQRCode(int(req.Price))

	// png, err := qrcode.Encode(reqQrCode, qrcode.Medium, 256)
	// if err != nil {
	// 	log.Printf("Error: Failed to create qrcode file: %s", err.Error())
	// 	// return nil, errors.New("error: failed to create qrcode file")
	// 	fileUrl = `https://i1.sndcdn.com/artworks-x8zI2HVC2pnkK7F5-4xKLyA-t1080x1080.jpg`
	// }

	// destination := fmt.Sprintf(u.uploadPath + utils.RandFileName())

	// fileUrl, err = gcpfile.UploadFile(u.cfg, u.cl, pctx, destination, png)
	// if err != nil {
	// 	log.Printf("Error: Upload file failed: %s", err.Error())
	// 	fileUrl = `https://i1.sndcdn.com/artworks-x8zI2HVC2pnkK7F5-4xKLyA-t1080x1080.jpg`
	// 	// return nil, errors.New("error: failed to upload file")
	// }
	fileUrl := string("https://storage.googleapis.com/ticket-shop-bucket/ticket-image/2409ec_1739993038586")

	orderNumber := utils.RandomString()

	if err := u.paymentRepo.AddTicketToCustomer(pctx, cfg, &payment.AddCustomerTicket{
		CustomerId:    req.CustomerId,
		OrderNumber:   orderNumber,
		Date:          req.Date,
		MovieName:     req.MovieName,
		MovieId:       req.MovieId,
		MovieDate:     req.MovieDate,
		MovieShowTime: req.MovieShowTime,
		PosterImage:   req.PosterImage,
		TicketUrl:     fileUrl,
		SeatNo:        req.SeatNo,
		Quantity:      req.Quantity,
		Price:         req.Price,
	}); err != nil {
		fmt.Printf("Error: Failed to add ticket %s", err.Error())
		return nil, errors.New("error: failed to add ticket")
	}

	return &payment.BuyticketRes{
		TransactionId: orderNumber,
		Url:           fileUrl,
	}, nil
}

// func (c *paymentUseCase) UploadFileTest(file multipart.File, object string) error {
// 	ctx := context.Background()

// 	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
// 	defer cancel()

// 	var png []byte
// 	png, err := qrcode.Encode("https://photos.app.goo.gl/pkN35vFQhc6DRXqQ6", qrcode.Medium, 256)
// 	if err != nil {
// 		log.Printf("Error: Failed to create qrcode file:%s", err.Error())
// 		return errors.New("error:failed to create qrcode file")
// 	}

// 	// buff := bytes.NewBuffer(png)

// 	gcpfile.UploadFile(c.cfg, c.cl, ctx, c.uploadPath+object, png)

// 	// Upload an object with storage.Writer.

// 	// wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
// 	// if _, err := io.Copy(wc, buff); err != nil {
// 	// 	fmt.Printf("Error:Failed to Upload File io.Copy: %s", err.Error())
// 	// 	return fmt.Errorf("io.Copy: %v", err)
// 	// }
// 	// if err := wc.Close(); err != nil {
// 	// 	fmt.Printf("Error:Failed to Upload File wc.Close: %s", err.Error())
// 	// 	return fmt.Errorf("Writer.Close: %v", err)
// 	// }

// 	return nil
// }
