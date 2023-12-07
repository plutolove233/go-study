package transaction

import (
	"block/constcoe"
	"block/utils"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

type Transaction struct {
	ID      []byte // the hash value of transaction, also can be seen as the id of transaction
	Inputs  []*TxInput
	Outputs []*TxOutput
}

func (tx *Transaction) TxHash() []byte {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	utils.Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	return hash[:]
}

func (tx *Transaction) SetID() {
	tx.ID = tx.TxHash()
}

func NewBaseTx(toaddress []byte) *Transaction {
	txIn := &TxInput{
		TxID:        []byte{},
		OutIdx:      -1,
		FromAddress: []byte{},
	}
	txOut := &TxOutput{
		constcoe.InitCoin,
		toaddress,
	}
	return &Transaction{
		[]byte("This is the base transaction!"),
		[]*TxInput{txIn},
		[]*TxOutput{txOut},
	}
}

func (tx *Transaction) IsBaseTx() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].OutIdx == -1
}
