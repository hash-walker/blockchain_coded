package blockchain

import (
	"time"
	"fmt"
	"math/rand/v2"
)

type Transactions struct{
	From string
	To string
	Amount int
	Timestamp time.Time
}

func InitTransaction() (Transactions,error){
	
	Users := makeUsers()

	fromUser := getRandomUser(Users, User{Name: ""})
	
	transactionAmount := rand.IntN(10)

	if transactionAmount <= 0{
		return Transactions{}, fmt.Errorf("no balance for the transaction")
	}


	fromUser.Amount -= transactionAmount
	toUser := getRandomUser(Users, fromUser)

	toUser.Amount += transactionAmount

	trx := Transactions{
		From: fromUser.Name,
		To: toUser.Name,
		Amount: transactionAmount,
		Timestamp: time.Now(),
	}

	return trx, nil
}