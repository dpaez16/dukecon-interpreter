package tests

import (
	"dukecon/lexer"
	"dukecon/token"
	"testing"
)

func TestGamevar(t *testing.T) {
	input := `
		gamevar TEMP0 1

		setvar TEMP0 2
		setvar TEMP0 -3

		addvar TEMP0 3
		subvar TEMP0 0
	`

	tests := []token.Token{
		{token.GAMEVAR, "gamevar"},
		{token.IDENT, "TEMP0"},
		{token.INT, "1"},

		{token.SETVAR, "setvar"},
		{token.IDENT, "TEMP0"},
		{token.INT, "2"},

		{token.SETVAR, "setvar"},
		{token.IDENT, "TEMP0"},
		{token.INT, "-3"},

		{token.ADDVAR, "addvar"},
		{token.IDENT, "TEMP0"},
		{token.INT, "3"},

		{token.SUBVAR, "subvar"},
		{token.IDENT, "TEMP0"},
		{token.INT, "0"},
	}

	evaluateTestCases(input, tests, t)
}

func TestIfStatements(t *testing.T) {
	input := `
		gamevar TEMP0 1

		ifvare TEMP0 0
		{
			nullop
		}
		else
		{
			setvar TEMP0 0
		}
	`

	tests := []token.Token{
		{token.GAMEVAR, "gamevar"},
		{token.IDENT, "TEMP0"},
		{token.INT, "1"},

		{token.IFVARE, "ifvare"},
		{token.IDENT, "TEMP0"},
		{token.INT, "0"},

		{token.LCURLY, "{"},
		{token.NULLOP, "nullop"},
		{token.RCURLY, "}"},
		{token.ELSE, "else"},
		{token.LCURLY, "{"},
		{token.SETVAR, "setvar"},
		{token.IDENT, "TEMP0"},
		{token.INT, "0"},
		{token.RCURLY, "}"},
	}

	evaluateTestCases(input, tests, t)
}

func TestReadFunction(t *testing.T) {
	input := `
		defstate foo
			setvar TEMP0 10
		ends

		state foo
	`

	tests := []token.Token{
		{token.DEFFUNCTION, "defstate"},
		{token.IDENT, "foo"},
		{token.SETVAR, "setvar"},
		{token.IDENT, "TEMP0"},
		{token.INT, "10"},
		{token.ENDFUNCTION, "ends"},

		{token.CALLFUNCTION, "state"},
		{token.IDENT, "foo"},
	}

	evaluateTestCases(input, tests, t)
}

func TestReadingNumbers(t *testing.T) {
	input := `
		-123
		12-2
		14
		0
	`

	tests := []token.Token{
		{token.INT, "-123"},
		{token.ILLEGAL, "12-2"},
		{token.INT, "14"},
		{token.INT, "0"},
	}

	evaluateTestCases(input, tests, t)
}

func TestReadingIdentifiers(t *testing.T) {
	input := `
		TEMP0
		0TEMP
		TEMP_2
	`

	tests := []token.Token{
		{token.IDENT, "TEMP0"},
		{token.ILLEGAL, "0TEMP"},
		{token.IDENT, "TEMP_2"},
	}

	evaluateTestCases(input, tests, t)
}

func evaluateTestCases(input string, tests []token.Token, t *testing.T) {
	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.Type {
			t.Fatalf("tests[%d] - tokentype wrong: expected=%q, got=%q",
				i, tt.Type, tok.Type)
		}

		if tok.Literal != tt.Literal {
			t.Fatalf("tests[%d] - literal wrong: expected=%q, got=%q",
				i, tt.Literal, tok.Literal)
		}
	}
}
