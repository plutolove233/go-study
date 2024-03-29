// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"io"
)

// Injectors from wire.go:

func InitOptionStruct(ctx context.Context, msg []Message, w io.Writer, r io.Reader) (*Greeter, error) {
	options := &Options{
		Message: msg,
		Writer:  w,
		Reader:  r,
	}
	greeter, err := NewGreeter(ctx, options)
	if err != nil {
		return nil, err
	}
	return greeter, nil
}
