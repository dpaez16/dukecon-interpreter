package parser

import (
	"dukecon/ast"
	"dukecon/token"
)

// gamevar <identifier> <value>

func (p *Parser) parseGamevarDeclareStatement() *ast.GamevarDeclareStatement {
	stmt := &ast.GamevarDeclareStatement{Token: p.currToken}

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
