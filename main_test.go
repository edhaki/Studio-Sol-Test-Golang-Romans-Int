package main

// para rodar o arquivo so digitar no terminal go test, vai verificar se ha erros.
import (
	"prova/romanos"
	"testing"
)

type romanValue struct {
	Roman string
	Value int
}

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		in   string
		want romanValue
	}{
		{"AXXBLX", romanValue{"LX", 60}},
		{"AIIVCHXL", romanValue{"C", 100}},
		{"AXVBXXKCLX", romanValue{"CLX", 160}},
		{"ABDLXIV", romanValue{"DLXIV", 564}},
	}

	for _, c := range cases {
		var roman, value = romanos.RomanNumerals(c.in)
		got := romanValue{roman, value}
		if got.Roman != c.want.Roman || got.Value != c.want.Value {
			t.Errorf("romanNumerals(%q)\ngot => {Roman: %q, Value: %d}\nwant => {Roman: %q, Value: %d}", c.in, got.Roman, got.Value, c.want.Roman, c.want.Value)
		}
	}
}
