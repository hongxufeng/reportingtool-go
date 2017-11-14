package function

import (
	"crypto/md5"
	"fmt"
)

// 根据字符串求取md5值
func Md5String(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func Md5UperString(s string) string {
	return fmt.Sprintf("%X", md5.Sum([]byte(s)))
}
