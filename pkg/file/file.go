package gcpfile

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
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

func UploadFile(cfg *config.Config, client *storage.Client, pctx context.Context, destination string, data []byte) (string, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*50)
	defer cancel()

	buff := bytes.NewBuffer(data)

	wc := client.Bucket(cfg.Gcp.BucketName).Object(destination).NewWriter(ctx)
	if _, err := io.Copy(wc, buff); err != nil {
		fmt.Printf("Error:Failed to Upload File io.Copy: %s", err.Error())
		return "", errors.New("faile to used io.copy")
	}
	if err := wc.Close(); err != nil {
		fmt.Printf("Error:Failed to Upload File wc.Close: %s", err.Error())
		return "", errors.New("faile to closed writer:")
	}

	if err := makePublic(cfg, ctx, client, destination); err != nil {
		fmt.Printf("Error:Faile to Make File Public: %s", err.Error())
		return "", errors.New("error: failed to make file public")
	}

	urlFile := fmt.Sprintf("https://storage.googleapis.com/%s/%s", cfg.Gcp.BucketName, destination)
	// urlFile := fmt.Sprint("https://img-cdn.pixlr.com/image-generator/history/65bb506dcb310754719cf81f/ede935de-1138-4f66-8ed7-44bd16efc709/medium.webp")

	// urlFile = "https://img-cdn.pixlr.com/image-generator/history/65bb506dcb310754719cf81f/ede935de-1138-4f66-8ed7-44bd16efc709/medium.webp"

	return urlFile, nil

}

func makePublic(cfg *config.Config, ctx context.Context, client *storage.Client, destination string) error {

	acl := client.Bucket(cfg.Gcp.BucketName).Object(destination).ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return fmt.Errorf("ACLHandle.Set: %w", err)
	}
	// fmt.Printf("Blob %v is now publicly accessible.\n", destination)
	return nil
}
