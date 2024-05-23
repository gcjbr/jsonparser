package lexer

type TokenType string
type Token struct {
	Type  TokenType
	Value string
}

const (
	EOF TokenType = "EOF"

	OPEN_CURLY_BRACE  TokenType = "OPEN_CURLY_BRACE"
	CLOSE_CURLY_BRACE TokenType = "CLOSE_CURLY_BRACE"
)
