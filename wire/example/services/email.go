// Package services
/*
@Coding : utf-8
@Time : 2023/2/9 22:18
@Author : yizhigopher
@Software : GoLand
*/
package services

import (
	"github.com/google/wire"
	"io"
	"log"
)

type MailService interface {
	Send()
}

type MailConfig struct {
	to      string
	from    string
	content string
}

type MailSender struct {
}

func NewMailSender(m *MailConfig) *MailSender {
	return &MailSender{}
}

func (m *MailSender) Send() {
	log.Println("send email")
}

var MailProvider = wire.NewSet(NewMailSender, wire.Bind(new(MailService), new(*MailSender)))

type Options struct {
	Message []string
	Writer  io.Writer
}

type Greeter struct {
	msg []string
	w   io.Writer
}

func NewGreeter(opts *Options) (*Greeter, error) {
	return &Greeter{
		msg: opts.Message,
		w:   opts.Writer,
	}, nil
}

var GreeterProvider = wire.NewSet(NewGreeter, wire.Struct(new(Options), "*"))
