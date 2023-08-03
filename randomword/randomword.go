package randomword

import (
	"bufio"
	"math/rand"
	"os"
)

var dict struct {
	words []string
	len   int
}

// init reads /usr/share/dict/web2 and stores the words in dict.words
func init() {
	file, err := os.Open("/usr/share/dict/web2")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dict.words = append(dict.words, scanner.Text())
	}
	dict.len = len(dict.words)
}

// RandomWord returns a random word from /usr/share/dict/web2
func RandomWord() string {
	return dict.words[rand.Intn(dict.len)]
}
