package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	words := []string{"apple", "banana", "cherry", "date", "elderberry",
		"fig", "grape", "honeydew", "iceberg lettuce", "jackfruit",
		"kiwi", "lemon", "mango", "nectarine", "orange", "pineapple",
		"quince", "raspberry", "strawberry", "tomato", "ugli fruit",
		"victoria plum", "watermelon", "xigua", "yellow passionfruit", "zucchini"}

	rand.Seed(time.Now().UnixNano())

	scanner := bufio.NewScanner(os.Stdin)
	score := 0
	done := make(chan bool)

	go func() {
		for {
			index := rand.Intn(len(words))
			fmt.Println("Type:", words[index])
			for {
				scanner.Scan()
				input := scanner.Text()
				if input == words[index] {
					score++
					break
				}
				fmt.Println("Type:", words[index])
			}
		}
	}()

	go func() {
		time.Sleep(60000 * time.Millisecond)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Time's up!")
		fmt.Println("Score:", score)
	}
}
