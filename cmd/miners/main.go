package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"time"
	"encoding/json"
	"github.com/hash-walker/blockchain_coded/internal/blockchain"
)

import (
    
)

func prettyPrintBlock(b blockchain.Block) {
    data, err := json.MarshalIndent(b, "", "  ")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(string(data))
}


type MemPool struct {
	Trx []blockchain.Transactions
}

func trxMemPool() MemPool {
	pool := MemPool{}

	for i := 0; i < 10; i++ {

		trx, err := blockchain.InitTransaction()

		if err != nil {
			fmt.Println(fmt.Errorf("trx is invalid"))
		} else {
			pool.Trx = append(pool.Trx, trx)
		}
	}

	return pool
}

func selectTrx(pool MemPool) []blockchain.Transactions {

	trx := make([]blockchain.Transactions,3)

	for i := 0; i < 3; i++ {
		randIdx := rand.IntN(3)
		value := pool.Trx[randIdx]
		trx[i] = value
		valudIdx := slices.Index(pool.Trx, value)
		pool.Trx = slices.Delete(pool.Trx, valudIdx, valudIdx+1)
	}

	return trx
}

func Miner (){
	poolIn := make(chan MemPool, 1)

	go func() {
		poolIn <- trxMemPool()
	}()

	pool := <-poolIn	

	trx := selectTrx(pool)

	Newblockchain := blockchain.Blockchain{
		Difficulty: 2,
	}

	genesisBlock := Newblockchain.NewBlockchain()
	Newblockchain.AddBlock(genesisBlock)

	block := blockchain.Block{
		Index:     genesisBlock.Index + 1,
		Timestamp: time.Now(),
		Nonce:     rand.IntN(100),
		Trx:       trx,
		PrevHash:  genesisBlock.Hash,
	}

	verifiedBlock := blockchain.MinerBlock(Newblockchain.Difficulty, block)

	Newblockchain.AddBlock(verifiedBlock)

	for _, block := range Newblockchain.Blocks {
		prettyPrintBlock(block)
	}
}

func main() {

	go func(){
		Miner()
	}()

	time.Sleep(2 * time.Second)

	go func(){
		Miner()
	}()

	time.Sleep(2 * time.Second)
}


