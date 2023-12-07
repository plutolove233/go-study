package pow

import (
	"block/transaction"
	"block/utils"
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	TimeStamp    time.Time
	Nonce        int64  // PoW
	Target       []byte // PoW
	Hash         []byte
	PrevHash     []byte
	Transactions []*transaction.Transaction
}

func (b *Block) GetTransactionHash() []byte {
	txIDs := make([][]byte, 0)
	for _, tx := range b.Transactions {
		txIDs = append(txIDs, tx.ID)
	}
	summary := bytes.Join(txIDs, []byte{})
	return summary
}

func (b *Block) SetHash() {
	information := bytes.Join([][]byte{
		utils.ToHexInt(b.TimeStamp.Unix()),
		utils.ToHexInt(int64(b.Nonce)),
		b.Target,
		b.PrevHash,
		b.GetTransactionHash(),
	}, []byte{})
	hash := sha256.Sum256(information)
	b.Hash = hash[:]
}

func NewBlock(prevHash []byte, txs []*transaction.Transaction) *Block {
	block := &Block{
		TimeStamp:    time.Now(),
		PrevHash:     prevHash,
		Transactions: txs,
	}
	block.Target = GetTarget()
	block.Nonce = block.FindNonce()
	block.SetHash()
	return block
}

func newGenesisBlock() *Block {
	tx := transaction.NewBaseTx([]byte("Orion Liu"))
	return NewBlock([]byte{}, []*transaction.Transaction{tx})
}
