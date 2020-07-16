package main

import (
	"fmt"
	"math/rand"
	"time"
)

var random_char = "!@#$%^&*()_+"

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 123))
	}
	return string(bytes)
}

func random_pool_str(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = random_char[rand.Intn(len(random_char))]
	}
	return string(bytes)
}

func main() {
	min, max := 1, 2000
	char_count := 10

	rand.Seed(time.Now().UnixNano()) //must
	fmt.Printf("Random number in range %d and %d is %d", min, max, randomInt(min, max))
	fmt.Println()
	fmt.Printf("Random string with %d char is %s", char_count, randomString(char_count))
	fmt.Println()
	fmt.Printf("Random Pool with %d char is %s", char_count, random_pool_str(char_count))
}
