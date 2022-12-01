package main

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// Check MD5不可逆，之可通过check函数来比对数据
func Check(content, encrypted string) bool {
	return strings.EqualFold(MD5Encode(content), encrypted)
}

func main() {
	strTest := "I love the world"
	println(MD5Encode(strTest))
}
