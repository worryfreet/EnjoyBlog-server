package encrypt

import (
	"fmt"
	"testing"
)

func TestRsaPubEncode(t *testing.T) {
	encode, err := RsaPubEncode("123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(encode)
	fmt.Println(RsaPriDecode(encode))
}
