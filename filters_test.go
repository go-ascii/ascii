package ascii

import (
	"testing"

	assert "gopkg.in/ktnyt/assert.v1"
)

func filterTest(set []byte, filter Filter) assert.F {
	cases := make([]assert.F, 128)
	for i := range cases {
		b := byte(i)
		v, e := filter(b), Contains(set)(b)
		cases[i] = assert.Equal(v, e)
	}
	return assert.All(cases...)
}

var testCases = []struct {
	name   string
	set    []byte
	filter Filter
}{
	{"Null", Null, IsNull},
	{"Graphic", Graphic, IsGraphic},
	{"Control", Control, IsControl},
	{"Space", Space, IsSpace},
	{"Quote", Quote, IsQuote},
	{"Upper", Upper, IsUpper},
	{"Lower", Lower, IsLower},
	{"Letter", Letter, IsLetter},
	{"Digit", Digit, IsDigit},
	{"Binary", Binary, IsBinary},
	{"Octal", Octal, IsOctal},
	{"Hex", Hex, IsHex},
	{"Latin", Latin, IsLatin},
	{"Snake", Snake, IsSnake},
	{"ASCII", ASCII, IsASCII},
	{"Byte", Byte, IsByte},
}

func TestFilters(t *testing.T) {
	for _, tt := range testCases {
		for i := 0; i < 128; i++ {
			b := byte(i)
			act, exp := tt.filter(b), Contains(tt.set)(b)
			if act != exp {
				t.Errorf("%[1]s(0x%[2]x) != Is%[1]s(0x%[2]x)", tt.name, b)
			}
		}
	}
}
