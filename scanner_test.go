package andrea_test

import (
	"github.com/CarolinaRomero33/z"
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok andrea.Token
		lit string
	}{
		{s: ``, tok: andrea.EOF},
		{s: `#`, tok: andrea.ILLEGAL, lit: `#`},
		{s: ` `, tok: andrea.WS, lit: " "},
		{s: "\t", tok: andrea.WS, lit: "\t"},
		{s: "\r", tok: andrea.WS, lit: "\r"},

		// Misc characters
		{s: `*`, tok: andrea.ASTERISK, lit: "*"},
		{s: `:`, tok: andrea.DosPunt, lit: ":"},
		{s: `(`, tok: andrea.ParDer, lit: "("},
		{s: `)`, tok: andrea.ParIsq, lit: ")"},
		{s: `;`, tok: andrea.PuntCom, lit: ";"},
		{s: `=`, tok: andrea.Asignaicon, lit: "="},
		// Identifiers
		{s: `foo`, tok: andrea.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: andrea.IDENT, lit: `Zx12_3U_`},

		{s: `zsen`, tok: andrea.zsen, lit: "zsen"},
		{s: `zcos`, tok: andrea.zcos, lit: "zcos"},
		{s: `ztan`, tok: andrea.ztan, lit: "ztan"},
		
		
		
	}

	for i, tt := range tests {
		s := andrea.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}

	}
}
