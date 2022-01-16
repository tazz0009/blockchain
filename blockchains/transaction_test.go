package blockchains_test

import (
	"fmt"
	"testing"

	"github.com/tazz0009/go-blockchain/blockchains"
)

func TestTransaction_CoinbaseTx(t *testing.T) {
	transaction := blockchains.CoinbaseTx("Jack", "")
	fmt.Println(transaction)
	if transaction.IsCoinbase() != true {
		t.Error()
	}
}
