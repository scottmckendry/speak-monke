package token

import (
	"testing"
)

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenType
	}{
		{"fn", FUNCTION},
		{"let", LET},
		{"true", TRUE},
		{"false", FALSE},
		{"if", IF},
		{"else", ELSE},
		{"return", RETURN},
	}

	for _, tt := range tests {
		if got := LookupIdent(tt.input); got != tt.expected {
			t.Errorf("LookupIdent(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}

	if got := LookupIdent("foobar"); got != IDENT {
		t.Errorf("LookupIdent(\"foobar\") = %q, want %q", got, IDENT)
	}
}
