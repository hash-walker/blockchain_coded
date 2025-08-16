package blockchain

import (
	"fmt"
	"strings"
	"math/rand/v2"
)

func MinerBlock(difficulty int, block Block) Block{
	var diff string
	for i:=0; i<difficulty; i++{
		diff += "0"
	}

	hashed := false
	iterations := 0
	var hash string

	for !hashed{
		hash = CalculateHash(block)
		hashed = strings.HasPrefix(hash, diff)
		block.Nonce = rand.IntN(1000)
		iterations++

		fmt.Printf("Hash: %v\n", hash)
	}

	block.Hash = hash
	return block
}