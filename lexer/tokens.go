package lexer

import (
	"fmt"
)

type TokenType int

const (
	EOF TokenType = iota
	OPEN_CURLY_BRACE
	CLOSE_CURLY_BRACE
	OPEN_SQUARE_BRACE
	CLOSE_SQUARE_BRACE
	OPEN_PARENTHESIS
	CLOSE_PARENTHESIS
	COLON
	COMMA
	STRING
	NUMBER
	TRUE
	FALSE
	NULL
	IDENTIFIER
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	return t.Value
}

func (t Token) Debug() {
	if t.isOneOf(STRING, IDENTIFIER, NUMBER) {
		fmt.Printf("%s (%s)\n", TokenTypeString(t.Type), t.Value)
	} else {
		fmt.Printf("%s ()\n", TokenTypeString(t.Type))
	}
}

func (t Token) isOneOf(expected ...TokenType) bool {
	for _, e := range expected {
		if t.Type == e {
			return true
		}
	}
	return false
}

func NewToken(t TokenType, v string) Token {
	return Token{Type: t, Value: v}
}

func TokenTypeString(t TokenType) string {
	switch t {
	case EOF:
		return "EOF"
	case OPEN_CURLY_BRACE:
		return "OPEN_CURLY_BRACE"
	case CLOSE_CURLY_BRACE:
		return "CLOSE_CURLY_BRACE"
	case OPEN_SQUARE_BRACE:
		return "OPEN_SQUARE_BRACE"
	case CLOSE_SQUARE_BRACE:
		return "CLOSE_SQUARE_BRACE"
	case OPEN_PARENTHESIS:
		return "OPEN_PARENTHESIS"
	case CLOSE_PARENTHESIS:
		return "CLOSE_PARENTHESIS"
	case COLON:
		return "COLON"
	case COMMA:
		return "COMMA"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case NULL:
		return "NULL"
	case IDENTIFIER:
		return "IDENTIFIER"
	}
	return "UNKNOWN"
}
