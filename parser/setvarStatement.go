package parser

import (
	"dukecon/ast"
	"dukecon/token"
)

// setvar <identifier> <value>

func (p *Parser) parseSetvarStatement() *ast.SetvarStatement {
	stmt := &ast.SetvarStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeek(token.INT) {
		return nil
	}

	stmt.Value = p.parseIntegerLiteral()

	return stmt
}
