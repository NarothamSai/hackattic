package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		txt := scanner.Text()
		if len(txt) == 0 {
			break
		}
		dec := decodeBase64String(txt)
		fmt.Println(dec)
	}

}

func decodeBase64String(base64Str string) string {
	dec, _ := base64.StdEncoding.DecodeString(base64Str)

	return string(dec[:])
}
