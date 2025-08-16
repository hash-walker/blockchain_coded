package blockchain

import (
	"math/rand/v2"
	"slices"
)

type User struct{
	Name string
	Amount int
}

func makeUsers() []User{
	users := []string{"Hamza", "Talha", "Okasha"}
	Users := make([]User, 3)
	for i := 0; i<len(users); i++{
		Users = append(Users, User{
			Name: users[i],
			Amount: rand.IntN(100),
		})
	}

	return Users
}

func getRandomUser (users []User, fromUser User) User{
	randomInt := rand.IntN(len(users))

	if fromUser.Name != "" {	
		userIndex := slices.Index(users, fromUser)

		for userIndex == randomInt{
			randomInt = rand.IntN(len(users))
		}

	}

	return users[randomInt]
}