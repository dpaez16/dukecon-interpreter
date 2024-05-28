package ast

import (
	"bytes"
	"dukecon/token"
)

// gamevar <identifier> <value>

type GamevarDeclareStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (gds *GamevarDeclareStatement) statementNode() {}

func (gds *GamevarDeclareStatement) TokenLiteral() string {
	return gds.Token.Literal
}

func (gds *GamevarDeclareStatement) String() string {
	var out bytes.Buffer

	out.WriteString(gds.TokenLiteral() + " ")
	out.WriteString(gds.Name.String() + " ")
	out.WriteString(gds.Value.String())

	return out.String()
}
