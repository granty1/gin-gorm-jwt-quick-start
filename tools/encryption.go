package tools

import (
	"crypto/md5"
	"fmt"
)

//Md5 return a hash value
func Md5(content string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}
