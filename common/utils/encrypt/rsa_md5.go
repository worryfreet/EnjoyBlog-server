package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/wenzhenxi/gorsa"
)

// RsaPubEncode 公钥加密
func RsaPubEncode(str string) (value string) {
	value, _ = gorsa.PublicEncrypt(str, string(publicKey))
	return
}

// RsaPriDecode 私钥解密
func RsaPriDecode(str string) (value string) {
	value, _ = gorsa.PriKeyDecrypt(str, string(privateKey))
	return
}

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum([]byte("worryfreet_EnjoyBlog")))
}
