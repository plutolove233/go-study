package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		println(err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		println(err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		println(err)
		return
	}

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"log",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		println(err)
		return
	}

	log.Printf("[x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
