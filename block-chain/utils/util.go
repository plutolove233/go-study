package utils

import (
	"bytes"
	"encoding/binary"
)

func ToHexInt(t int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, t)
	Handle(err)
	return buff.Bytes()
}
