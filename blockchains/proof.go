package blockchains

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

const Difficulty = 18

type ProofWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofWork{b, target}

	return pow
}

func (pow *ProofWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransations(),
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{})
	return data
}

func (pow *ProofWork) Run() (int, []byte) {
	var initHash big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		initHash.SetBytes(hash[:])
		if initHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]

}

func (pow *ProofWork) Validate() bool {
	var initHash big.Int
	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	initHash.SetBytes(hash[:])
	return initHash.Cmp(pow.Target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	CheckError(err)
	return buff.Bytes()
}
