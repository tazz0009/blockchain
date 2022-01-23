package goblockchain_test

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	blocks "github.com/tazz0009/go-blockchain/blockchains"
)

type sendForm struct {
	from   string
	to     string
	amount int
}

type testCase struct {
	sendArr []sendForm
}

func TestMain(t *testing.T) {
	var sendArr []sendForm
	sendArr = append(sendArr, sendForm{"17CauH6yWr86k5rS18rNHQSnWjwJ1BkUcE", "1MrEFL1w8JbyKGXW2SveS5ZmQw7MMwEGzo", 50})
	sendArr = append(sendArr, sendForm{"1MrEFL1w8JbyKGXW2SveS5ZmQw7MMwEGzo", "1E9sKGBDNQuxt8nXdkKNoPWkgeaTX3KVZX", 20})
	sendArr = append(sendArr, sendForm{"17CauH6yWr86k5rS18rNHQSnWjwJ1BkUcE", "1E9sKGBDNQuxt8nXdkKNoPWkgeaTX3KVZX", 30})
	// createWallet(t)

	// listAddress(t)

	// createBlockChain(t)

	printChain(t)

	// printBalanceOnlyOne(t)

	// send(t, sendArr[0])
	// send(t, sendArr[1])
	// send(t, sendArr[2])

}

func send(t *testing.T, form sendForm) {
	if !blocks.ValidateAddress(form.to) {
		log.Panic("Address is not Valid")
	}
	if !blocks.ValidateAddress(form.from) {
		log.Panic("Address is not Valid")
	}
	chain := blocks.ContinueBlockChain(form.from)
	defer chain.Database.Close()

	tx := blocks.NewTransaction(form.from, form.to, form.amount, chain)
	chain.AddBlock([]*blocks.Transaction{tx})
	fmt.Println("Success!")
}

func printBalanceOnlyOne(t *testing.T) {
	address := "17CauH6yWr86k5rS18rNHQSnWjwJ1BkUcE"
	if !blocks.ValidateAddress(address) {
		log.Panic("Address is not Valid")
	}
	chain := blocks.ContinueBlockChain(address)
	defer chain.Database.Close()

	balance := 0
	pubKeyHash := blocks.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := chain.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of %s: %d\n", address, balance)
}

func printChain(t *testing.T) {

	fmt.Println("printChain!!!")
	chain := blocks.ContinueBlockChain("")
	defer chain.Database.Close()
	iter := chain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		pow := blocks.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func createBlockChain(t *testing.T) {
	address := "17CauH6yWr86k5rS18rNHQSnWjwJ1BkUcE"

	if !blocks.ValidateAddress(address) {
		log.Panic("Address is not Valid")
	}
	chain := blocks.InitBlockChain(address)
	chain.Database.Close()
	fmt.Println("Finished!")
}

func listAddress(t *testing.T) {
	wallets, _ := blocks.CreateWallets()
	addresses := wallets.GetAllAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}

func createWallet(t *testing.T) {
	fmt.Println("createWallet!!!")
	wallets, _ := blocks.CreateWallets()
	address := wallets.AddWallet()
	wallets.SaveFile()

	fmt.Printf("New address is: %s\n", address)
}

func createBlockchain(t *testing.T) {
	fmt.Println("createBlockchain!!!")
}
