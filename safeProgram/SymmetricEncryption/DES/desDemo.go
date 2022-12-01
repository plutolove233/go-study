package main

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
	"fmt"
)

var key = []byte("qoxiu7Z1") // 只能为8位

func ZeroPadding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)
	return append(cipherText, padText...)
}

func ZeroUnPadding(padText []byte) []byte {
	return bytes.TrimFunc(padText, func(r rune) bool {
		return r == rune(0)
	})
}

func DesEncrypt(plainText string) (string, error) {
	src := []byte(plainText)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("need a multiple of the block size")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func DesDecrypt(cipherText string) (string, error) {
	src, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("input not full blocks")
	}

	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}

func main() {
	text := "你哈"
	cipherText, err := DesEncrypt(text)
	if err != nil {
		fmt.Println("加密失败, err=", err)
		return
	}
	println("密文为:", cipherText)
	plainText, err := DesDecrypt(cipherText)
	if err != nil {
		fmt.Println("解密失败, err=", err)
		return
	}
	println("明文为:", plainText)
}
