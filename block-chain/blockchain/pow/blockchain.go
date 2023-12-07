package pow

import (
	"block/transaction"
	"block/utils"
	"encoding/hex"
	"fmt"
	"sync"
)

type BlockChain struct {
	rwmu   sync.RWMutex
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	bc := &BlockChain{}
	bc.Blocks = append(bc.Blocks, newGenesisBlock())
	return bc
}

func (bc *BlockChain) AddBlock(txs []*transaction.Transaction) {
	b := NewBlock(bc.Blocks[len(bc.Blocks)-1].Hash, txs)
	bc.rwmu.Lock()
	defer bc.rwmu.Unlock()
	bc.Blocks = append(bc.Blocks, b)
}

func (bc *BlockChain) FindUnspentTransactions(address []byte) []*transaction.Transaction {
	unspentTxs := make([]*transaction.Transaction, 0)
	spentTxs := make(map[string][]int)

	for idx := len(bc.Blocks) - 1; idx >= 0; idx-- {
		b := bc.Blocks[idx]
		for _, tx := range b.Transactions {
			txID := hex.EncodeToString(tx.ID)

		IterOutput:
			for outIdx, out := range tx.Outputs {
				if spentTxs[txID] != nil {
					for _, spentOut := range spentTxs[txID] {
						if spentOut == outIdx {
							continue IterOutput
						}
					}
				}

				if out.ToAddressRight(address) {
					unspentTxs = append(unspentTxs, tx)
				}
			}
			if !tx.IsBaseTx() {
				for _, in := range tx.Inputs {
					if in.FromAddressRight(address) {
						inTxID := hex.EncodeToString(in.TxID)
						spentTxs[inTxID] = append(spentTxs[inTxID], in.OutIdx)
					}
				}
			}
		}
	}
	return unspentTxs
}

func (bc *BlockChain) FindUTXOs(address []byte) (int, map[string]int) {
	unspentOuts := make(map[string]int)
	unspentTxs := bc.FindUnspentTransactions(address)

	accumulated := 0

Work:
	for _, tx := range unspentTxs {
		txID := hex.EncodeToString(tx.ID)
		for outIdx, out := range tx.Outputs {
			if out.ToAddressRight(address) {
				accumulated += out.Value
				unspentOuts[txID] = outIdx
				continue Work
			}
		}
	}

	return accumulated, unspentOuts
}

func (bc *BlockChain) FindSpendableOutputs(address []byte, amount int) (int, map[string]int) {
	unspentOuts := make(map[string]int)
	unspentTxs := bc.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTxs {
		txID := hex.EncodeToString(tx.ID)
		for outIdx, out := range tx.Outputs {
			if out.ToAddressRight(address) && accumulated < amount {
				accumulated += out.Value
				unspentOuts[txID] = outIdx
				if accumulated >= amount {
					break Work
				}
				continue Work
			}
		}
	}
	return accumulated, unspentOuts
}

func (bc *BlockChain) CreateTransaction(from, to []byte, amount int) (*transaction.Transaction, bool) {
	var inputs []*transaction.TxInput
	var outputs []*transaction.TxOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
	if acc < amount {
		fmt.Println("Not enough coins!")
		return nil, false
	}

	for txid, outidx := range validOutputs {
		txID, err := hex.DecodeString(txid)
		utils.Handle(err)
		input := &transaction.TxInput{
			TxID: txID,
			OutIdx: outidx,
			FromAddress: from,
		}
		inputs = append(inputs, input)
	}

	outputs = append(outputs, &transaction.TxOutput{
		Value:     amount,
		ToAddress: to,
	})

	if acc > amount {
		outputs = append(outputs, &transaction.TxOutput{
			Value: acc - amount,
			ToAddress: from,
		})
	}
	tx := &transaction.Transaction{
		ID: nil,
		Inputs: inputs,
		Outputs: outputs,
	}
	tx.SetID()
	return tx, true
}
