package encryutil

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5String 对string md5
func Md5String(src string) string {
	return Md5([]byte(src))
}

// Md5 md5
func Md5(src []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(src)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
