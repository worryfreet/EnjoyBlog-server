package encrypt

import (
	"fmt"
	"testing"
)

func TestRsaPubEncode(t *testing.T) {
	encode := RsaPubEncode("123456")
	fmt.Println(encode)
	fmt.Println(RsaPriDecode(encode))
}
