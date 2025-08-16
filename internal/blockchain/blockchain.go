package blockchain

import (
	"encoding/hex"
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	"github.com/google/uuid"
)


type Blockchain struct{
	Blocks []Block
	Difficulty int
}

func (b *Blockchain) NewBlockchain() Block{

	fmt.Print("Starting to generate genesis block....\n")

	transactions := make([]Transactions, 1)

	transactions = append(transactions, Transactions{
		From: uuid.NewString(),
		To: uuid.NewString(),
		Amount: rand.IntN(100),
		Timestamp: time.Now(),
	})

	block := Block{
		Index: 0,
		Timestamp: time.Now(),
		Nonce: rand.IntN(10),
		Trx: transactions,
		PrevHash: hex.EncodeToString([]byte("0000000000000000")),
	}
	
	var difficulty string
	for i:=0; i<b.Difficulty; i++{
		difficulty += "0"
	}

	hashed := false
	iterations := 0
	var hash string
	for !hashed{
		hash = CalculateHash(block)
		hashed = strings.HasPrefix(hash, difficulty)
		block.Nonce = rand.IntN(1000)
		iterations++

		fmt.Printf("Hash: %v\n", hash)
	}

	block.Hash = hash

	fmt.Printf("There is our genesis block\n")

	return block
}


func (b *Blockchain) AddBlock(block Block){
	b.Blocks = append(b.Blocks, block)
}
