package main

import (
	"context"
	"github.com/google/wire"
	"io"
)

type Message string

type Options struct {
	Message []Message
	Writer  io.Writer
	Reader  io.Reader
}

type Greeter struct {
}

func NewGreeter(c context.Context, options *Options) (*Greeter, error) {
	return nil, nil
}

var GreeterSet = wire.NewSet(wire.Struct(new(Options), "*"), NewGreeter)

func main() {
	//InitOptionStruct(context.Background(), []Message{"hello", "World"}, io.Writer, io.Reader)
}
