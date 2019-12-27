package main

/*

You have 50 bitcoins to distribute to 10 users: Matthew, Sarah, Augustus,
Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth The coins will be distributed
based on the vowels contained in each name where:
a: 1 coin e: 1 coin i: 2 coins o: 3 coins u: 4 coins
and a user can’t get more than 10 coins. Print a map with each user’s name
and the amount of coins distributed.

*/

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	coins := 50.
	fmt.Printf("%#v \nCoins left: %2.0f", distributeCoins([]string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}, &coins), coins)
}

func distributeCoins(users []string, totalCoins *float64) map[string]float64 {
	const maxCoinsPerUser = 10.
	distribution := make(map[string]float64, len(users))
	for _, userName := range users {
		userCoins := math.Min(maxCoinsPerUser, getUserCoins(userName))
		*totalCoins -= userCoins
		distribution[userName] = userCoins
	}
	return distribution
}

func getUserCoins(name string) float64 {
	total := 0.
	vowelsDist := map[string]float64{"a": 1, "e": 1, "i": 2, "o": 3, "u": 4}

	for _, v := range name {
		total += vowelsDist[strings.ToLower(string(v))]
	}
	return total
}
