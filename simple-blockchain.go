// You can edit this code!
// Click here and start typing.

// Asim Kamran

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type blockchain struct {
	Blockchain      []*Block
	TransactionPool []Transaction
}

func (bc *blockchain) NewBlock(nonce int) *Block {

	if bc.TransactionPool != nil {
		var st = Block{}
		st.TimeStamp = strconv.FormatInt(time.Now().UnixNano(), 10)
		for i := range bc.TransactionPool {
			st.Transaction = append(st.Transaction, bc.TransactionPool[i])
		}
		bc.TransactionPool = nil
		st.Nonce = nonce

		if len(bc.Blockchain) == 0 {
			st.PreviousHash = ""
		} else {
			st.PreviousHash = bc.Blockchain[len(bc.Blockchain)-1].BlockHash
		}
		out, err := json.Marshal(st.Transaction)
		trans := hex.EncodeToString(out[:])

		if err != nil {
			panic(err)
		}

		st.BlockHash = CalculateHash(trans + strconv.Itoa(st.Nonce) + st.TimeStamp + st.PreviousHash)

		return &st
	} else {
		fmt.Printf("No transaction to add .\n\n")
		return nil
	}

}

func CalculateHash(hashh string) string {

	return fmt.Sprintf("%x", sha256.Sum256([]byte(hashh)))
}

func (bc *blockchain) AddTransaction(sender string, recipient string, value float32) {
	trans := NewTransaction(sender, recipient, value)
	bc.TransactionPool = append(bc.TransactionPool, *trans)
}

func (bc *blockchain) PrintBlocks() {

	for i := range bc.Blockchain {
		fmt.Printf("%s Block No %d %s\n", strings.Repeat("*", 25), i, strings.Repeat("*", 25))
		fmt.Printf("Time Stamp: %s\n", bc.Blockchain[i].TimeStamp)
		fmt.Printf("nonce: %d\n", bc.Blockchain[i].Nonce)
		fmt.Printf("Previous Hash: %s\n", bc.Blockchain[i].PreviousHash)
		fmt.Printf("Block Hash: %s\n\n", bc.Blockchain[i].BlockHash)
		fmt.Printf("%s Transactions %s\n", strings.Repeat("*", 22), strings.Repeat("*", 21))

		val, err := json.MarshalIndent(bc.Blockchain[i].Transaction, "", "    ")
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n\n", val)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 58))

}

type Block struct {
	Transaction  []Transaction
	TimeStamp    string
	Nonce        int
	PreviousHash string
	BlockHash    string
}

type Transaction struct {
	TransactionId              string
	SenderBlockchainAddress    string
	RecipientBlockchainAddress string
	Value                      float32
}

func (t *Transaction) SetTransactionId() {
	sum := sha256.Sum256([]byte(t.SenderBlockchainAddress + t.RecipientBlockchainAddress + strconv.FormatFloat(float64(t.Value), 'E', -1, 32)))
	t.TransactionId = hex.EncodeToString(sum[:])
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	trans := new(Transaction)
	trans.SenderBlockchainAddress = sender
	trans.RecipientBlockchainAddress = recipient
	trans.Value = value
	trans.SetTransactionId()
	return trans
}

/*
func VerifyChain(bk *blockchain) bool {
	var st = ""
	for i := 0; i < len(bk.list); i++ {

		st = CalculateHash(strconv.Itoa(bk.list[i].nonce) + bk.list[i].prevHash)

		if st != bk.list[i].currHash {
			fmt.Printf("Tempered Block no %d\n", i)
			return false

		}

	}

	return true

}*/

/*func ChangeBlock(bk *blockchain, nonce int, transaction string) {

	for i := 0; i < len(bk.list); i++ {
		if nonce == bk.list[i].nonce {

			bk.list[i].transaction = transaction
			fmt.Println("change done ")
			return
		}
	}

	fmt.Println("block not found!")

}*/

func main() {
	blockchain := new(blockchain)
	blockchain.AddTransaction("Dell", "Hp", 12.5)
	blockchain.AddTransaction("KFC", "McDonalds", 3.3)
	blockchain.AddTransaction("Fast", "Nust", 3.5)
	blockchain.AddTransaction("Toyota", "Honda", 5.2)
	blk := blockchain.NewBlock(987654321)
	blockchain.Blockchain = append(blockchain.Blockchain, blk)

	blockchain.AddTransaction("Linex", "Windows", 1.5)
	blockchain.AddTransaction("Suzuki", "Yamaha", 3.7)
	blk = blockchain.NewBlock(12345678)
	blockchain.Blockchain = append(blockchain.Blockchain, blk)
	blockchain.PrintBlocks()

	/*blockchain.addblock(10, "Dell to HP 5000")
	blockchain.addblock(11, "KFC to MCdonalds 2500")
	blockchain.addblock(12, "Fast to Nust 150")
	blockchain.addblock(13, "Toyota to Honda 1500")
	blockchain.addblock(14, "Linux to windows 1200")
	fmt.Println("All Block Added")
	ListBlocks(blockchain)
	fmt.Println("")
	fmt.Println("")*/

	/*fmt.Println("Block Nonce 12 Changed")
	ChangeBlock(blockchain, 12, "Fast to Cust 150")
	ListBlocks(blockchain)
	*/

	/*fmt.Println("")
	fmt.Println("")/*/

	/*fmt.Println("Add New Block ")
	blockchain.addblock(15, "Asim to asim 2200")
	fmt.Println("Chain Verification ")
	VerifyChain(blockchain)
	fmt.Println("Hashes Relinking")
	blockhash(blockchain)
	fmt.Println("Chain Verification ")
	VerifyChain(blockchain)
	fmt.Println("Add New Block ")
	blockchain.addblock(15, "Asim to asim 2200")
	/*/

}
