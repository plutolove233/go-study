package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func loadTLSConfig(caFile string) *tls.Config {
	// load tls config
	var tlsConfig tls.Config
	tlsConfig.InsecureSkipVerify = false
	if caFile != "" {
		certpool := x509.NewCertPool()
		ca, err := os.ReadFile(caFile)
		if err != nil {
			log.Fatal(err.Error())
		}
		certpool.AppendCertsFromPEM(ca)
		tlsConfig.RootCAs = certpool
	}
	return &tlsConfig
}

var (
	messagePublicHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
		fmt.Printf("Received message: %s from topic %s\n", m.Payload(), m.Topic())
	}

	connectHandler mqtt.OnConnectHandler = func(c mqtt.Client) {
		fmt.Println("Connected")
	}
	connectLostHandler mqtt.ConnectionLostHandler = func(c mqtt.Client, err error) {
		fmt.Printf("Connect loss: %v\n", err)
	}
)

func main() {
	tlsConfig := loadTLSConfig("./emqxsl-ca.crt")
	protocol := "ssl"
	broker := "x026b664.ala.us-east-1.emqxsl.com"
	// broker := "broker.emqx.io"
	port := 8883
	options := mqtt.NewClientOptions()
	// 连接远程服务
	options.AddBroker(fmt.Sprintf("%s://%s:%d", protocol, broker, port))
	options.SetClientID("go_mqtt_demo_x026b664")
	options.SetUsername("mqtt_demo")
	options.SetPassword("123456")
	options.SetKeepAlive(time.Second * 60)
	options.SetTLSConfig(tlsConfig)

	// 挂载相关的处理函数
	options.SetDefaultPublishHandler(messagePublicHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	publish(client)
	client.Disconnect(250)
}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("subscribed to topic: %s\n", topic)
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("test-message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}
