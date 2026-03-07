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

func (u *paymentUseCase) BuyTicketConsumer(pctx context.Context, cfg *config.Config, topic string, key string, resCh chan<- *payment.RollBackReserveSeatRes, readyCh chan<- struct{}) {
	// ใช้ low-level kafka.Conn แทน kafka.Reader
	// DialLeader + Seek จะ block จนกว่าเชื่อมต่อสำเร็จจริงๆ
	conn, err := queue.KafkaConnConsumer(u.cfg, topic)
	if err != nil {
		log.Printf("Error creating consumer connection: %s", err.Error())
		close(resCh)
		readyCh <- struct{}{}
		return
	}
	defer conn.Close()

	// ถึงตรงนี้ connection พร้อมแล้ว 100% → ส่งสัญญาณ ready
	readyCh <- struct{}{}

	type msgResult struct {
		data *payment.RollBackReserveSeatRes
		err  error
	}
	msgCh := make(chan msgResult, 1)

	go func() {
		for {
			// ReadBatch อ่าน message จาก partition โดยตรง (ไม่ผ่าน consumer group)
			batch := conn.ReadBatch(1, 10e6) // minBytes=1, maxBytes=10MB

			for {
				msg, err := batch.ReadMessage()
				if err != nil {
					break // end of batch
				}

				if string(msg.Key) == key {
					data := new(payment.RollBackReserveSeatRes)
					if err := json.Unmarshal(msg.Value, data); err != nil {
						fmt.Printf("Error: Unmarshal error %s\n", err.Error())
						msgCh <- msgResult{err: err}
						batch.Close()
						return
					}
					msgCh <- msgResult{data: data}
					batch.Close()
					return
				}
			}
			batch.Close()
		}
	}()

	select {
	case <-pctx.Done():
		log.Println("BuyTicketConsumer context cancelled")
		close(resCh)
		return
	case result := <-msgCh:
		if result.err != nil {
			log.Printf("Error reading message: %s", result.err.Error())
			close(resCh)
			return
		}
		resCh <- result.data
		close(resCh)
		return
	}
}

func (u *paymentUseCase) BuyTicket(pctx context.Context, cfg *config.Config, req *payment.MovieBuyReq) (*payment.BuyticketRes, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	stage1 := new(payment.RollBackReserveSeatRes)

	resCh := make(chan *payment.RollBackReserveSeatRes, 1)
	readyCh := make(chan struct{}, 1)

	go u.BuyTicketConsumer(ctx, cfg, "reserve-seat-res", "payment", resCh, readyCh)

	// รอให้ consumer เชื่อมต่อ Kafka สำเร็จก่อน ค่อยส่ง ReserveSeat
	select {
	case <-readyCh:
		log.Println("Consumer is ready, sending ReserveSeat message...")
	case <-ctx.Done():
		return nil, errors.New("timeout waiting for consumer to be ready")
	}

	req.CustomerId = strings.TrimPrefix(req.CustomerId, "customer:")

	if err := u.paymentRepo.ReserveSeat(ctx, cfg, &payment.ReserveSeatReq{
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
	// case <-time.After(time.Second * 30):
	case <-ctx.Done():
		u.paymentRepo.RollBackReserveSeat(ctx, cfg, &payment.RollBackReservedSeatReq{
			MovieId: req.MovieId,
			SeatNo:  req.SeatNo,
		})
		fmt.Println("Timeout waiting for rollback response")
		return nil, errors.New("timeout waiting for rollback response")
	}

	if stage1.Error != "" {
		fmt.Println("stage1.Error", stage1.Error)
		u.paymentRepo.RollBackReserveSeat(ctx, cfg, &payment.RollBackReservedSeatReq{
			MovieId: req.MovieId,
			SeatNo:  req.SeatNo,
		})

		return nil, errors.New("error: failed to reserve seat")
	}

	if err := u.CheckOutWithCreditCard(&payment.CheckOutWithCreditCard{Token: req.Token, Price: req.Price}); err != nil {
		u.paymentRepo.RollBackReserveSeat(ctx, cfg, &payment.RollBackReservedSeatReq{
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

	// fileUrl, err = gcpfile.UploadFile(u.cfg, u.cl, ctx, destination, png)
	// if err != nil {
	// 	log.Printf("Error: Upload file failed: %s", err.Error())
	// 	fileUrl = `https://i1.sndcdn.com/artworks-x8zI2HVC2pnkK7F5-4xKLyA-t1080x1080.jpg`
	// 	// return nil, errors.New("error: failed to upload file")
	// }
	fileUrl := string("https://storage.googleapis.com/ticket-shop-bucket/ticket-image/2409ec_1739993038586")

	orderNumber := utils.RandomString()

	if err := u.paymentRepo.AddTicketToCustomer(ctx, cfg, &payment.AddCustomerTicket{
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
