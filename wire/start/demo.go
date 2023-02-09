package main

import "fmt"

type Message struct {
	msg string
}

type Greeter struct {
	message Message
}

type Event struct {
	greeter Greeter
}

func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

func NewGreeter(m Message) Greeter {
	return Greeter{
		message: m,
	}
}

func NewEvent(g Greeter) Event {
	return Event{greeter: g}
}

func (g *Greeter) Greet() Message {
	return g.message
}

func (e *Event) Start() {
	msg := e.greeter.Greet()
	fmt.Println(msg)
}

func main() {
	//message := NewMessage("hello world")
	//greeter := NewGreeter(message)
	//event := NewEvent(greeter)
	event := initEvent("xxx")

	event.Start()
}
