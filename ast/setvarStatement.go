package ast

import (
	"bytes"
	"dukecon/token"
)

// setvar <identifier> <value>

type SetvarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (svs *SetvarStatement) statementNode() {}

func (svs *SetvarStatement) TokenLiteral() string {
	return svs.Token.Literal
}

func (svs *SetvarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(svs.TokenLiteral() + " ")
	out.WriteString(svs.Name.String() + " ")
	out.WriteString(svs.Value.String())

	return out.String()
}
