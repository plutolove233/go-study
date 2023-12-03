package pow

import (
	"block/utils"
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	TimeStamp time.Time
	Nonce     int64  // PoW
	Target    []byte // PoW
	Hash      []byte
	PrevHash  []byte
	Data      []byte
}

func (b *Block) SetHash() {
	information := bytes.Join([][]byte{
		utils.ToHexInt(b.TimeStamp.Unix()), 
		utils.ToHexInt(int64(b.Nonce)),
		b.Target,
		b.PrevHash, 
		b.Data,
	}, []byte{})
	hash := sha256.Sum256(information)
	b.Hash = hash[:]
}

func NewBlock(prevHash []byte, data []byte) *Block {
	block := &Block{
		TimeStamp: time.Now(),
		PrevHash:  prevHash,
		Data:      data,
	}
	block.Target = GetTarget()
	block.Nonce = block.FindNonce()
	block.SetHash()
	return block
}

func newGenesisBlock() *Block {
	return NewBlock([]byte{}, []byte("Hello Block Chain"))
}
