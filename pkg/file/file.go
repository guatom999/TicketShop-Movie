package file

import (
	"context"
	"errors"
	"log"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"github.com/guatom999/TicketShop-Movie/config"
)

type (
	ClientUpLoader struct {
		cl         *storage.Client
		projectID  string
		bucketName string
		uploadPath string
	}

	FileReq struct {
		File        *multipart.FileHeader `form:"file"`
		Destination string                `form:"destination"`
		Extension   string                `form:"extension"`
		FileName    string                `form:"filename"`
	}

	FileRes struct {
		FileName string `json:"filename"`
		Url      string `json:"url"`
	}
)

func UploadFile(cfg *config.Config, data []byte) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	newCli, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Error: Failed To Connect GCP")
		return errors.New("error: failed to connect gcp")
	}

	cli := &ClientUpLoader{
		cl:         newCli,
		projectID:  "ticket-shop-427608",
		bucketName: "ticket-shop-bucket",
		// uploadPath: "ticket-image",
		uploadPath: "ticket-image2/",
	}

	_ = cli.cl.Bucket("ticket-shop-bucket").Object(cli.uploadPath + "test-01").NewWriter(ctx)
	// wc.ContentType()

	// buff := bytes.NewBuffer(b)

	// if _, err := io.Copy(wc, data); err != nil {
	// 	return fmt.Println("io.Copy: %v", err)
	// }
	// if err := wc.Close(); err != nil {
	// 	return fmt.Println("Writer.Close: %v", err)
	// }

	// buff := bytes.NewBuffer(data)

	// if _, err = io.Copy(wc, buff); err != nil {
	// 	log.Fatalf("Error: failed to upload file: %s", err.Error())
	// 	return errors.New("error: failed to upload file ")
	// }

	return nil
}
