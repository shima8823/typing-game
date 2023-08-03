package main

import (
	"bufio"
	"fmt"
	"github.com/shima8823/typing-game/randomword"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	score := 0

	go func() {
		for {
			word := randomword.RandomWord()
			fmt.Println("Type:", word)
			for {
				scanner.Scan()
				input := scanner.Text()
				if input == word {
					score++
					break
				}
				fmt.Println("Type:", word)
			}
		}
	}()

	select {
	case <-time.After(20 * time.Second):
		fmt.Println("Time's up!")
		fmt.Println("Score:", score)
	}
}
