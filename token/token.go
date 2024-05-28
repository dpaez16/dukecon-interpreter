package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// comments
	LINECOMMENT       = "//"
	BEGINBLOCKCOMMENT = "/*"
	ENDBLOCKCOMMENT   = "*/"

	// identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// code blocks, switch statements, gamearray access, struct access
	LCURLY = "{"
	RCURLY = "}"
	LBRACK = "["
	RBRACK = "]"

	// switch statements (case)
	COLON = ":"

	// keywords for functions
	DEFFUNCTION  = "defstate"
	ENDFUNCTION  = "ends"
	CALLFUNCTION = "state"

	// keywords for variables
	GAMEVAR   = "gamevar"
	GAMEARRAY = "gamearray"

	// gamevar operations
	SETVAR = "setvar"
	ADDVAR = "addvar"
	SUBVAR = "subvar"

	// keywords for if statements
	IFVARE    = "ifvare"
	IFVARVARE = "ifvarvare"
	IFVARN    = "ifvarn"
	IFVARVARN = "ifvarvarn"

	// flow control
	ELSE   = "else"
	NULLOP = "nullop"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[TokenType]struct{}{
	GAMEVAR:      {},
	SETVAR:       {},
	ADDVAR:       {},
	SUBVAR:       {},
	DEFFUNCTION:  {},
	ENDFUNCTION:  {},
	CALLFUNCTION: {},
	IFVARE:       {},
	NULLOP:       {},
	ELSE:         {},
}

func LookupIdent(ident string) TokenType {
	tokenType := TokenType(ident)
	if _, ok := keywords[tokenType]; ok {
		return tokenType
	}

	return IDENT
}
