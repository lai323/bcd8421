
#  BCD 8421 codec

[BCD](https://zh.wikipedia.org/zh-cn/%E4%BA%8C%E9%80%B2%E7%A2%BC%E5%8D%81%E9%80%B2%E6%95%B8) 8421 编解码

输入输出为数字字符串，没有数字大小限制

### example

```go
package main

import (
	"fmt"

	"github.com/lai323/bcd8421"
)

func main() {

	bytes, _ := bcd8421.EncodeFromStr("18749009700", 10)
	fmt.Println(fmt.Sprintf("%#x", bytes)) // output: 0x00000000018749009700

	s, _ := bcd8421.DecodeToStr(bytes, true)
	fmt.Println(s) // output: "18749009700"

}

```