package blockchains

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

type Block struct {
	Hash         []byte
	Transactions []*Transaction
	PrevHash     []byte
	Nonce        int
}

func (b *Block) HashTransations() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, txs, prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	CheckError(err)
	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	CheckError(err)
	return &block
}

func CheckError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (b *Block) String() string {
	return fmt.Sprintf("Hash:%x\nTransactions:[%s]\nPrevHash:%x\nNonce:%d\n", b.Hash, b.Transactions, b.PrevHash, b.Nonce)
}
