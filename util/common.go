package util

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5加密字符串
func MD5Str(str string) string{
	md5 := md5.New()
	md5.Write([]byte(str))
	return hex.EncodeToString(md5.Sum(nil))
}
