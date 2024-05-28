package tests

import (
	"dukecon/ast"
	"dukecon/lexer"
	"dukecon/parser"
	"testing"
)

func TestGamevarDeclareStatements(t *testing.T) {
	input := `
		gamevar TEMP0 1
		gamevar TEMP1 -1
		gamevar TEMP2 0
		gamevar TEMP3 100
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	checkParserErrors(t, p)

	if len(program.Statements) != 4 {
		t.Fatalf("program.Statements does not contain 4 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
		expectedValue      int64
	}{
		{"TEMP0", 1},
		{"TEMP1", -1},
		{"TEMP2", 0},
		{"TEMP3", 100},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testGamevarDeclareStatement(t, stmt, tt.expectedIdentifier, tt.expectedValue) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}

func testGamevarDeclareStatement(t *testing.T, s ast.Statement, name string, value int64) bool {
	if s.TokenLiteral() != "gamevar" {
		t.Errorf("s.TokenLiteral not 'gamevar'. got=%q", s.TokenLiteral())
		return false
	}

	stmt, ok := s.(*ast.GamevarDeclareStatement)
	if !ok {
		t.Errorf("s not *ast.GamevarDeclareStatement. got=%T", s)
		return false
	}

	if stmt.Name.Value != name {
		t.Errorf("stmt.Name.Value not '%s'. got=%s", name, stmt.Name.Value)
		return false
	}

	if stmt.Name.TokenLiteral() != name {
		t.Errorf("stmt.Name.TokenLiteral() not '%s'. got=%s", name, stmt.Name.TokenLiteral())
		return false
	}

	literal, ok := stmt.Value.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("stmt.Value not *ast.IntegerLiteral. got=%T", stmt.Value)
		return false
	}

	if literal.Value != value {
		t.Errorf("literal.Value not %d. got=%d", value, literal.Value)
		return false
	}

	return true
}
