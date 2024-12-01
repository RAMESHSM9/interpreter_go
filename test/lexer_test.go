package test

import (
	"testing"

	"github.com/RAMESHSM9/interpreter_go/src/token"

	"github.com/RAMESHSM9/interpreter_go/src/lexer"
)

func TestNextToken(t *testing.T) {

	input := `+(){},;`

	tests := []struct {
		expectedTokenType token.TokenType
		expectedLiteral   string
	}{
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACES, "{"},
		{token.RBRACES, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for i, tokenType := range tests {

		tok := l.NextToken()

		if tok.Type != tokenType.expectedTokenType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tokenType.expectedTokenType, tok.Type)
		}

		if tok.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tokenType.expectedLiteral, tok.Literal)
		}

	}

}
