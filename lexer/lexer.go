package lexer

import (
	"fmt"
	"regexp"
)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type regexHandler func(lex *lexer, regex *regexp.Regexp)

type lexer struct {
	Tokens   []Token
	patterns []regexPattern
	source   string
	position int
}

func Tokenize(source string) []Token {
	lex := createLexer(source)

	for !lex.at_eof() {
		matched := false
		remainder := lex.remainder()
		//fmt.Printf("Current remainder: %s\n", remainder)

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(remainder)
			//fmt.Printf("Checking pattern: %v, loc: %v\n", pattern.regex.String(), loc)

			if loc != nil && loc[0] == 0 {
				//fmt.Println("Match found, handling...")
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			fmt.Printf("No match found for remainder: %s at position: %d\n", remainder, lex.position)
			panic("Lexer error at " + remainder)
		}
	}
	lex.push(NewToken(EOF, "EOF"))

	return lex.Tokens
}

func (lex *lexer) advance(n int) {
	lex.position += n
}

func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) at() byte {
	return lex.source[lex.position]
}

func (lex *lexer) remainder() string {
	return lex.source[lex.position:]
}

func (lex *lexer) at_eof() bool {
	return lex.position >= len(lex.source)
}

func defaultHandler(tokenType TokenType, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		// Advance lexer position by the length of the match
		lex.advance(len(value))
		// Push new token to lexer's token list
		lex.push(NewToken(tokenType, value))
	}
}

func createLexer(source string) *lexer {
	return &lexer{
		position: 0,
		source:   source,
		Tokens:   make([]Token, 0),
		patterns: []regexPattern{
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY_BRACE, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY_BRACE, "}")},
			{regexp.MustCompile(`\s+`), skipHandler},
		},
	}
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
	lex.advance(len(regex.FindString(lex.remainder())))
}
