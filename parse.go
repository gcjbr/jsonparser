package main

import (
	"fmt"
	"os"

	"github.com/gcjbr/jsonparser/lexer"
)

func main() {
	bytes, _ := os.ReadFile("tests/step1/valid.json")
	l := lexer.NewLexer(string(bytes))

	for {
		t := l.NextToken()
		if t.Type == lexer.EOF {
			break
		}
		fmt.Println(t.Type)
	}

}
