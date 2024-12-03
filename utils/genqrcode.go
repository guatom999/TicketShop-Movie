package utils

import (
	"fmt"
	"strconv"

	"github.com/sigurn/crc16"
)

func GenQRCode(price int) string {

	price = price / 100
	PrompPayKey := crc16.Params{0x1021, 0xFFFF, false, false, 0x0000, 0x0000, "CRC-16/XMODEM"}
	table := crc16.MakeTable(PrompPayKey)

	lenPrice := len(strconv.Itoa(price)) + 3
	priceQrString := fmt.Sprintf("%s.00", strconv.Itoa(price))

	result := fmt.Sprintf("00020101021129370016A000000677010111011300668887874005802TH5303764540%v%v", lenPrice, priceQrString)

	lastFoutDig := crc16.Checksum([]byte(fmt.Sprintf("%v6304", result)), table)

	return fmt.Sprintf("%v6304%X", result, lastFoutDig)
}
