package main

import (
	"fmt"

	"github.com/lai323/bcd8421"
)

func main() {
	var bytestr []string
	bytes, _ := bcd8421.EncodeFromStr("18749009700", 10)

	for _, b := range bytes {
		bytestr = append(bytestr, fmt.Sprintf("%#x", b))
	}
	fmt.Println(bytestr) // output: [0x0 0x0 0x0 0x0 0x1 0x87 0x49 0x0 0x97 0x0]

	s, _ := bcd8421.DecodeToStr(bytes)
	fmt.Println(s) // output: "18749009700"
}
