package pow

import (
	"block/constcoe"
	"block/utils"
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

func GetTarget() []byte {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-constcoe.Difficulty))
	return target.Bytes()
}

func (b *Block) GetBase4Nonce(nonce int64) []byte {
	data := bytes.Join([][]byte{
		utils.ToHexInt(b.TimeStamp.Unix()),
		utils.ToHexInt(int64(nonce)),
		b.Target,
		b.PrevHash,
		b.GetTransactionHash(),
	}, []byte{})
	return data
}

func (b *Block) FindNonce() int64 {
	var intHash big.Int
	var intTarget big.Int
	var hash [32]byte
	var nonce int64
	nonce = 0
	intTarget.SetBytes(b.Target)

	for nonce < math.MaxInt64 {
		data := b.GetBase4Nonce(nonce)
		hash = sha256.Sum256(data)
		intHash.SetBytes(hash[:])
		if intHash.Cmp(&intTarget) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce
}

func (b *Block) ValidatePoW() bool {
	var intHash big.Int
	var intTarget big.Int
	var hash [32]byte
	intTarget.SetBytes(b.Target)
	data := b.GetBase4Nonce(b.Nonce)
	hash = sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(&intTarget) == -1
}