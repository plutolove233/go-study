package transaction

import "bytes"

type TxOutput struct {
	Value     int    // the value of transaction
	ToAddress []byte // the address of receiver
}

type TxInput struct {
	TxID        []byte // the former transaction id
	OutIdx      int    // the id of transaction's output
	FromAddress []byte // the address of sender
}

func (in *TxInput) FromAddressRight(address []byte) bool {
	return bytes.Equal(in.FromAddress, address)
}

func (out *TxOutput) ToAddressRight(address []byte) bool {
	return bytes.Equal(out.ToAddress, address)
}
