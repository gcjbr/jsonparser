package lexer

type Lexer struct {
	input        string
	position     int  // Points to the current char (ch)
	readPosition int  // Points to the next char to be read
	ch           byte // current char
}

// Read current char and advance the lexer to the next char

func (l *Lexer) readChar() {
	// If end of input, set ch to 0
	// ASCII for NUL
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() Token {
	var t Token

	switch l.ch {
	case '{':
		t.Type = OPEN_CURLY_BRACE
	case '}':
		t.Type = CLOSE_CURLY_BRACE
	case 0:
		t.Type = EOF
	default:
		t.Type = EOF
	}
	l.readChar()

	return t
}

func newToken(t TokenType, ch byte) Token {
	return Token{
		Type:  t,
		Value: string(ch),
	}
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()

	return l
}
