//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"io"
)

func InitOptionStruct(ctx context.Context, msg []Message, w io.Writer, r io.Reader) (*Greeter, error) {
	wire.Build(GreeterSet)
	return nil, nil
}
