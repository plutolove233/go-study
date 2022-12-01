package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		fmt.Println("connect to RabbitMQ failed, err: ", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("open a channel failed, err: ", err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue",
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println("declare a queue failed, err: ", err)
		return
	}

	body := bodyForm(os.Args)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		fmt.Println("publish a message failed, err: ", err)
		return
	}

	log.Printf(" [x] Sent %s", body)
}

func bodyForm(args []string) string {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
