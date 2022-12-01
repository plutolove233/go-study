package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

func RSAEncrypt(plainText []byte) ([]byte, error) {
	file, err := os.Open("./rsa_public.pem")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(buf)
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub := publicInterface.(*rsa.PublicKey)

	return rsa.EncryptPKCS1v15(rand.Reader, pub, plainText)
}

func RSADecrypt(cipherText []byte) ([]byte, error) {
	file, err := os.Open("./rsa_private.pem")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	_, err = file.Read(buf)

	block, _ := pem.Decode(buf)

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
}

func main() {
	data, err := RSAEncrypt([]byte("ifinjlefnijnwiuew93024234sldfneiuwfnieruwnirunv"))
	if err != nil {
		panic(err)
	}
	println(base64.StdEncoding.EncodeToString(data))
	original, err := RSADecrypt(data)
	if err != nil {
		panic(err)
	}
	println(string(original))
}
