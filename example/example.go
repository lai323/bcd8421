package main

import (
	"fmt"

	"github.com/lai323/bcd8421"
)

func main() {

	bytes, _ := bcd8421.EncodeFromStr("18749009700", 10)
	fmt.Println(fmt.Sprintf("%#x", bytes)) // output: 0x00000000018749009700

	s, _ := bcd8421.DecodeToStr(bytes)
	fmt.Println(s) // output: "18749009700"

}
