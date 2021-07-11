
#  BCD 8421 codec

[BCD](https://zh.wikipedia.org/zh-cn/%E4%BA%8C%E9%80%B2%E7%A2%BC%E5%8D%81%E9%80%B2%E6%95%B8) 8421 编解码

输入输出为数字字符串，没有数字大小限制

### example

```go
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

```