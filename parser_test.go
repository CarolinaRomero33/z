package andrea_test

import (
	"github.com/CarolinaRomero33/z"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)  

func TestParser_ParseStatement(t *testing.T) {
	f, _ := ioutil.ReadFile("C:/Users/Romero/Escritorio/iop.txt")
	z := string(f[:])


	
	var tests = []struct {
		s    string
		stmt *andrea.Statement
		err  string
	}{
		
		{
			s:    z,
			stmt: &andrea.Statement{
	
			},
		},

		
	}

	for i, tt := range tests {
		stmt, err := andrea.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}


func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
