/*
Pow is one way which need each node to guess the nonce.
Firstly the nonce is record in the head of Block.
Secondly we should calculate the hash (within the nonce), and the hash is lower than target difficulty
Thirdly, who guess the nonce, who can join the block chain
*/

package constcoe

/*
 constcoe is used to discribute the difficulty of PoW,
 which is the alogrithm of block chain
*/
const (
	Difficulty = 20
	InitCoin   = 1000
)
