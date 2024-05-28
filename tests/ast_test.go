package tests

import (
	"dukecon/ast"
	"dukecon/token"
	"testing"
)

func TestStr(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.GamevarDeclareStatement{
				Token: token.Token{Type: token.GAMEVAR, Literal: "gamevar"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "TEMP0"},
					Value: "TEMP0",
				},
				Value: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "123"},
					Value: 123,
				},
			},
		},
	}

	if program.String() != "gamevar TEMP0 123" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
