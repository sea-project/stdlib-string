package stdlib_string

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// GetRandString 随机生成N位字符串
func GetRandString(n int) string {
	mainBuff := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, mainBuff)
	if err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	return hex.EncodeToString(mainBuff)[:n]
}

// GetMd5String 生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

