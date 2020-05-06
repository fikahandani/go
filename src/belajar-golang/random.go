package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Mengenal Nilai Golang")
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random first value is ", randomValue)
	randomValue = randomWithRange(10, 50)
	fmt.Println("Random second value is ", randomValue)
	randomValue = randomWithRange(50, 100)
	fmt.Println("Random third value is ", randomValue)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}
