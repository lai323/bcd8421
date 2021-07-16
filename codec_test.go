package bcd8421

import (
	"reflect"
	"testing"
)

func TestStringNumberToBytes(t *testing.T) {
	should := []byte{0x00, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9}
	b, err := stringNumberToBytes("0123456789")
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(b, should) {
		t.Errorf("should be %#v", should)
	}
}

var testcases = []struct {
	number      string
	bytes       []byte
	bytesLength int
}{
	{
		number:      "907865438",
		bytes:       []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x09, 0x07, 0x86, 0x54, 0x38},
		bytesLength: 10,
	},
	{
		number:      "9007865438",
		bytes:       []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x90, 0x07, 0x86, 0x54, 0x38},
		bytesLength: 10,
	},
	{
		number:      "90007865438",
		bytes:       []byte{0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x07, 0x86, 0x54, 0x38},
		bytesLength: 10,
	},
	{
		number:      "900007865438",
		bytes:       []byte{0x00, 0x00, 0x00, 0x00, 0x90, 0x00, 0x07, 0x86, 0x54, 0x38},
		bytesLength: 10,
	},
	{
		number:      "9000007865438",
		bytes:       []byte{0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x07, 0x86, 0x54, 0x38},
		bytesLength: 10,
	},
	{
		number:      "3830",
		bytes:       []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x38, 0x30},
		bytesLength: 10,
	},
	{
		number:      "38300",
		bytes:       []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x83, 0x00},
		bytesLength: 10,
	},
	{
		number:      "383000",
		bytes:       []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x38, 0x30, 0x00},
		bytesLength: 10,
	},
	{
		number:      "3830000",
		bytes:       []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x83, 0x00, 0x00},
		bytesLength: 10,
	},
}

func TestEncodeFromStr(t *testing.T) {
	for _, tt := range testcases {
		ret, err := EncodeFromStr(tt.number, tt.bytesLength)
		if err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(ret, tt.bytes) {
			t.Errorf("EncodeFromStr(%#v) = %#v; should be %#v", tt.number, ret, tt.bytes)
		}
	}
}

func TestDecodeToStr(t *testing.T) {
	for _, tt := range testcases {
		n, err := DecodeToStr(tt.bytes)
		if err != nil {
			panic(err)
		}
		if n != tt.number {
			t.Errorf("DecodeToStr(%#v) = %s; should be %s", tt.bytes, n, tt.number)
		}
	}
}
