package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var key = []byte("hesisyizhigopher") // 密钥字符长度为16、24、32

func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padNum := blockSize - len(cipherText)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	cipherText = append(cipherText, ret...)
	return cipherText
}

func PKCS7UnPadding(originData []byte) []byte {
	length := len(originData)
	unpadding := int(originData[length-1])
	return originData[:(length - unpadding)]
}

func AesEncrypt(src string) (string, error) {
	srcByte := []byte(src)

	block, err := aes.NewCipher(key)
	if err != nil {
		return src, err
	}

	newSrcByte := PKCS7Padding(srcByte, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(newSrcByte))
	blockMode.CryptBlocks(dst, newSrcByte)

	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd, nil
}

func AesDecrypt(cipherText string) (string, error) {
	cipherByte, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return cipherText, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return cipherText, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])

	plainByte := make([]byte, len(cipherByte))

	blockMode.CryptBlocks(plainByte, cipherByte)
	plainByte = PKCS7UnPadding(plainByte)
	return string(plainByte), nil
}

func main() {
	text := "你好"
	cipherText, err := AesEncrypt(text)
	if err != nil {
		fmt.Println("加密失败, err=", err)
		return
	}
	println("密文为:", cipherText)
	plainText, err := AesDecrypt(cipherText)
	if err != nil {
		fmt.Println("解密失败, err=", err)
		return
	}
	println("明文为:", plainText)
}
