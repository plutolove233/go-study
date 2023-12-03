package pow

import "sync"

type BlockChain struct {
	rwmu   sync.RWMutex
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	bc := &BlockChain{}
	bc.Blocks = append(bc.Blocks, newGenesisBlock())
	return bc
}

func (bc *BlockChain) AddBlock(data string) {
	b := NewBlock(bc.Blocks[len(bc.Blocks)-1].Hash, []byte(data))
	bc.rwmu.Lock()
	defer bc.rwmu.Unlock()
	bc.Blocks = append(bc.Blocks, b)
}
