package blockchains_test

import (
	"fmt"
	"testing"

	"github.com/tazz0009/go-blockchain/blockchains"
)

type testCase struct {
	address    string
	sendTo     string
	sendAmount int
}

/*
 * Genesis Block 생성
 * tazz0009 : 100
 */
// func TestBlockChain_InitBlockChain(t *testing.T) {
// 	tc := testCase{"tazz0009", "", 0}
// 	blockchain := blockchains.InitBlockChain(tc.address)
// 	defer blockchain.Database.Close()
// }

/*
 * 1 Block 생성
 * tazz0009 : 30 -> tom0001
 */
/*
 * 2 Block 생성
 * tom0001 : 20 -> tazz0009
 */
/*
 * 3 Block 생성
 * tazz0009 : 50 -> tom0001
 */
/*
 * 4 Block 생성
 * tazz0009 : 10 -> tom0001
 */
/*
 * 5 Block 생성
 * tom0001 : 30 -> tazz0009
//  */
func TestBlockChain_send(t *testing.T) {
	tc := testCase{"tazz0009", "tom0001", 30}
	blockchain := blockchains.ContinueBlockChain(tc.address)
	defer blockchain.Database.Close()

	tx := blockchains.NewTransaction(tc.address, tc.sendTo, tc.sendAmount, blockchain)
	blockchain.AddBlock([]*blockchains.Transaction{tx})
}

func TestBlockChain_getBalance(t *testing.T) {
	tc := testCase{"tazz0009", "tom0001", 50}
	blockchain := blockchains.ContinueBlockChain(tc.address)
	defer blockchain.Database.Close()

	balance := 0
	UTXOs := blockchain.FindUTXO(tc.address)
	for _, out := range UTXOs {
		balance += out.Value
	}
	fmt.Printf("Balance of %s: %d\n", tc.address, balance)
}

/*
 * Print All blocks
 */
// func TestBlockChain_Print(t *testing.T) {
// 	const dbPath = "../tmp/blocks"
// 	var lastHash []byte
// 	db, err := badger.Open(badger.DefaultOptions(dbPath))
// 	blockchains.CheckError(err)

// 	err = db.View(func(txn *badger.Txn) error {
// 		item, err := txn.Get([]byte("lh"))
// 		blockchains.CheckError(err)
// 		err = item.Value(func(val []byte) error {
// 			lastHash = val
// 			fmt.Printf("lh:%x\n", lastHash)
// 			return nil
// 		})
// 		return err
// 	})
// 	blockchains.CheckError(err)

// 	blockchain := blockchains.BlockChain{lastHash, db}
// 	iter := blockchain.Iterator()
// 	for {
// 		block := iter.Next()
// 		fmt.Println(block)
// 		if len(block.PrevHash) == 0 {
// 			break
// 		}
// 	}
// }

func TestBlockChain_FindUnspentTransactions(t *testing.T) {
	tc := testCase{"tazz0009", "tom0001", 50}
	blockchain := blockchains.ContinueBlockChain(tc.address)
	defer blockchain.Database.Close()

	unspentTxs := blockchain.FindUnspentTransactions(tc.address)
	for _, tx := range unspentTxs {
		fmt.Println(tx.String())
	}
}
