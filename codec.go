package bcd8421

import (
	"bytes"
	"fmt"
	"strconv"
)

// bytesLength 是生成的 BCD 长度，如果 number 长度不足，在前面补 0
func EncodeFromStr(number string, bytesLength int) ([]byte, error) {
	var (
		numberBytes  []byte
		numberLength = len(number)
	)
	if bytesLength*2 < numberLength {
		return numberBytes, fmt.Errorf("invalid bytesLength")
	}

	nb, err := stringNumberToBytes(number)
	if err != nil {
		return numberBytes, err
	}
	if numberLength%2 == 1 {
		nb = append([]byte{0x00}, nb...)
	}
	if fill := bytesLength*2 - len(nb); fill != 0 {
		nb = append(bytes.Repeat([]byte{0x00}, fill), nb...)
	}

	// n1 左移 4 位，高位丢弃，低位补 0，得到 n3
	// n3 n2 进行或操作，即可将 n1 的低 4 位和 n2 的低 4 位压缩在一起
	for i := 0; i < len(nb); i += 2 {
		n1 := nb[i]
		n2 := nb[i+1]
		n3 := n1 << 4
		numberBytes = append(numberBytes, n3|n2)
		// fmt.Printf("%d, %d : %#b %#b %#b %#b\n", nb[i], nb[i+1], n1, n2, n3, n3|n2)
	}
	return numberBytes, nil
}

func DecodeToStr(src []byte) (string, error) {
	var (
		s          string
		foundFirst bool
	)

	for _, b := range src {
		if b == 0x00 && !foundFirst {
			continue
		}
		// 高 4 位直接是第一个数字
		n1 := b >> 4
		// 将低 4 位左移到高 4 位，在将高低 4 位调换位置得到 n2
		mask := b << 4
		n2 := mask<<4 | mask>>4
		// fmt.Printf("%#b n1:%#b %#b n2:%#b\n", b, n1, b<<4, n2)
		if n1 > 9 || n2 > 9 {
			return s, fmt.Errorf("invalid BCD 8421 bytes")
		}
		if n1 != 0x00 && !foundFirst {
			foundFirst = true
		}
		if n1 != 0x00 || foundFirst {
			s += strconv.Itoa(int(n1))
		}
		if n2 != 0x00 || foundFirst {
			s += strconv.Itoa(int(n2))
		}
	}
	return s, nil
}

// 将字符串数字逐个先转为数字，再转为 byte
func stringNumberToBytes(number string) ([]byte, error) {
	const fnAtoi = "stringNumberToBytes"

	var b []byte
	for _, ch := range []byte(number) {
		ch -= '0'
		if ch > 9 {
			return b, &strconv.NumError{Func: fnAtoi, Num: number, Err: strconv.ErrSyntax}
		}
		b = append(b, ch)
	}
	return b, nil
}
