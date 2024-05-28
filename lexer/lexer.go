package lexer

import (
	"dukecon/token"
	"regexp"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '{':
		tok = newToken(token.LCURLY, l.ch)
	case '}':
		tok = newToken(token.RCURLY, l.ch)
	case '[':
		tok = newToken(token.LBRACK, l.ch)
	case ']':
		tok = newToken(token.RBRACK, l.ch)
	case '/':
		if l.peekChar() == '/' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)

			tok = token.Token{Type: token.LINECOMMENT, Literal: literal}
		} else if l.peekChar() == '*' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)

			tok = token.Token{Type: token.BEGINBLOCKCOMMENT, Literal: literal}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case '*':
		if l.peekChar() == '/' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)

			tok = token.Token{Type: token.ENDBLOCKCOMMENT, Literal: literal}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()

			if isNumber(tok.Literal) {
				tok.Type = token.INT
			} else {
				tok.Type = token.LookupIdent(tok.Literal)
			}

			if tok.Type == token.IDENT && !isIdentifier(tok.Literal) {
				tok.Type = token.ILLEGAL
			}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

		return tok
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || isDigit(ch)
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == '-'
}

func isNumber(s string) bool {
	re := regexp.MustCompile(`^-?\d+$`)
	return re.MatchString(s)
}

func isIdentifier(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z_\d]+$`)
	return re.MatchString(s)
}
