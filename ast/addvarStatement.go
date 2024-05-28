package ast

import (
	"bytes"
	"dukecon/token"
)

// addvar <identifier> <value>

type AddvarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (avs *AddvarStatement) statementNode() {}

func (avs *AddvarStatement) TokenLiteral() string {
	return avs.Token.Literal
}

func (avs *AddvarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(avs.TokenLiteral() + " ")
	out.WriteString(avs.Name.String() + " ")
	out.WriteString(avs.Value.String())

	return out.String()
}
