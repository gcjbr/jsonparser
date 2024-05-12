package main

import (
	"os"

	"github.com/gcjbr/jsonparser/lexer"
)

func main() {
	bytes, _ := os.ReadFile("tests/step1/valid.json")
	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		token.Debug()
	}
}
