package raftrpc

import (
	"net/rpc"
)

type ClientEnd struct {
	connect bool   // this end-point is connected?
	Address string // this end-point's address
}

// send an RPC, wait for the reply.
// the return value indicates success; false means that
// no reply was received from the server.
func (e *ClientEnd) Call(svcMeth string, args interface{}, reply interface{}) bool {
	c, err := rpc.DialHTTP("tcp", e.Address)
	if err != nil {
		return false
	}
	err = c.Call(svcMeth, args, reply)
	if err != nil {
		return false
	}
	return true
}

func NewClientEnd(address string) *ClientEnd {
	return &ClientEnd{
		Address: address,
	}
}
