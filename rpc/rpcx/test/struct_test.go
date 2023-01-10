package test

import (
	"rpcx"
	"testing"
)

func TestStruct(t *testing.T) {
	x := &rpcx.Greeter{}
	y := new(rpcx.Greeter)
	t.Log(x, y, x == y)
}
