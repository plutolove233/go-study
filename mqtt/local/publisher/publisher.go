package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	protocol = "mqtt"
	address  = "localhost"
	port     = 1883
)

type Data struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewData(msg string) Data {
	return Data{
		Code: 200,
		Msg:  msg,
	}
}

func main() {
	opt := mqtt.NewClientOptions()
	broker := fmt.Sprintf("%s://%s:%d", protocol, address, port)
	// 可以连接多个broker组成的集群
	opt.AddBroker(broker)
	opt.SetClientID("mqttx_123456")

	// 设置遗嘱信息，遗嘱信息只会在出现意外情况时才会发送
	lastWillData := NewData("finished!")
	b, _ := json.Marshal(lastWillData)
	opt.SetBinaryWill("local/info", b, 0, false)
	opt.SetKeepAlive(time.Second * 1)
	/*
		if you use tls, or we don't suggest you set username and password

		opt.SetUsername("qqqq")
		opt.SetPassword("123456")
	*/

	client := mqtt.NewClient(opt)

	token := client.Connect()
	if token.Wait(); token.Error() != nil {
		fmt.Println("connect to local mqtt broker failed, error=", token.Error())
		return
	}
	for i := 0; i < 10; i++ {
		data := NewData(fmt.Sprintf("Hello: %d", i))
		b, err := json.Marshal(data)
		if err != nil {
			log.Fatalln("marshal failed, err=", err)
			continue
		}
		client.Publish("local/info", 0, false, b)
	}

	// 正常结束，不会发送遗嘱信息
	client.Disconnect(250)
}
