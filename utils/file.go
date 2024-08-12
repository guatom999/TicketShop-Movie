package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func RandFileName() string {
	fileName := fmt.Sprintf("%s_%v", uuid.NewString()[:6], time.Now().UnixMilli())

	return fileName
}
