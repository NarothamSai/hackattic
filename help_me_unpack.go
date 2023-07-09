package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	base64String := "5frVjHN3xPrWagAAB3Z6QjnO1HD4W3NAQHNb+HDUzjk="

	decoded := decodeBase64String(base64String)

	signedInt := getSignedIntegerFromBytes(decoded[0:4])
	unsignedInt := getUnsignedIntegerFromBytes(decoded[4:8])
	shortInt := getShortIntFromBytes(decoded[8:12])
	floatNum := float64(getFloatFromBytes(decoded[12:16]))
	doubleNum := getDoubleFromBytes(decoded[16:24])
	doubleBigEndianNum := getDoubleBigEndianFromBytes(decoded[24:32])

	fmt.Println("signedInt:", signedInt)
	fmt.Println("unsignedInt:", unsignedInt)
	fmt.Println("shortInt:", shortInt)
	fmt.Println("floatNum:", floatNum)
	fmt.Println("doubleNum:", doubleNum)
	fmt.Println("doubleBigEndianNum:", doubleBigEndianNum)
}

func decodeBase64String(base64Str string) []byte {
	dec, _ := base64.StdEncoding.DecodeString(base64Str)

	return dec
}

func getUnsignedIntegerFromBytes(bytesData []byte) uint32 {
	return binary.LittleEndian.Uint32(bytesData)
}

func getSignedIntegerFromBytes(bytesData []byte) int32 {
	return int32(binary.LittleEndian.Uint32(bytesData))
}

func getShortIntFromBytes(bytesData []byte) int16 {
	return int16(binary.LittleEndian.Uint16(bytesData))
}

func getFloatFromBytes(bytesData []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(bytesData))
}

func getDoubleFromBytes(bytesData []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(bytesData))
}

func getDoubleBigEndianFromBytes(bytesData []byte) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(bytesData))
}
