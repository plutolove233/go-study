package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const _crt = "../cert/localhost.pem"

var uri = "https://localhost:5555"

func main() {
	request2()
}

func request2() {
	client := &http.Client{
		Timeout: 1 * time.Hour,
	}
	//读取证书文件，正式环境无读取证书文件，因为本地测试是无法认证证书
	caCert, err := ioutil.ReadFile(_crt)
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	//tls协议配置，InsecureSkipVerify认证证书是否跳过
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
		//设置安全跳跃认证
		InsecureSkipVerify: true,
	}
	client.Transport = &http2.Transport{
		TLSClientConfig: tlsConfig,
	}

	urlObj, _ := url.Parse(uri)
	params := url.Values{}
	params.Add("message", "Hello H2")
	urlObj.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", urlObj.String(), nil)
	req.Header.Set("Extra", "123456")

	resp, err := client.Get(urlObj.String())
	fmt.Println(resp.StatusCode)
}
