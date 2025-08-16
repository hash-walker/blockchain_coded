package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Index int
	Timestamp time.Time
	Nonce int
	Trx []Transactions
	PrevHash string
	Hash string
}

func (b *Block) Transaction() string{
	record := make([]string, 4)
	for _,trx := range b.Trx {
		record = append(record, fmt.Sprintf(
		"%s->%s%d@%s",
		trx.From,
		trx.To,
		trx.Amount,
		trx.Timestamp.Format(time.RFC3339),
	))
	}

	return strings.Join(record, "|")
}

func CalculateHash (block Block) string{
	record := fmt.Sprintf(
		"%d%s%d%s%s",
		block.Index,
		block.PrevHash,
		block.Nonce,
		block.Timestamp.Format(time.RFC3339),
		block.Transaction(),
	)

	hash := sha256.Sum256([]byte(record))

	return hex.EncodeToString(hash[:])
}