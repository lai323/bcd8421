package bcd8421

import (
	"fmt"
	"testing"
)

func TestStringNumberToBytes(t *testing.T) {
	b, err := stringNumberToBytes("0123456789")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b)
}

func TestEncodeFromStr(t *testing.T) {
	bytes, err := EncodeFromStr("1234000580000", 10)
	if err != nil {
		panic(err)
	}
	for _, b := range bytes {
		fmt.Printf("%#x ", b)
	}
	fmt.Println()
	s, err := DecodeToStr(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
