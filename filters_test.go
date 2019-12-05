package ascii_test

import (
	"testing"

	"gopkg.in/ktnyt/ascii.v1"
	"gopkg.in/ktnyt/assert.v1"
)

func filterTest(set []byte, filter ascii.Filter) assert.F {
	cases := make([]assert.F, 128)
	for i := range cases {
		b := byte(i)
		v, e := filter(b), ascii.Contains(set)(b)
		cases[i] = assert.Equal(v, e)
	}
	return assert.All(cases...)
}

type testCaseType struct {
	Name   string
	Set    []byte
	Filter ascii.Filter
}

var testCases = []testCaseType{
	testCaseType{"Null", ascii.Null, ascii.IsNull},
	testCaseType{"Graphic", ascii.Graphic, ascii.IsGraphic},
	testCaseType{"Control", ascii.Control, ascii.IsControl},
	testCaseType{"Space", ascii.Space, ascii.IsSpace},
	testCaseType{"Quote", ascii.Quote, ascii.IsQuote},
	testCaseType{"Upper", ascii.Upper, ascii.IsUpper},
	testCaseType{"Lower", ascii.Lower, ascii.IsLower},
	testCaseType{"Letter", ascii.Letter, ascii.IsLetter},
	testCaseType{"Digit", ascii.Digit, ascii.IsDigit},
	testCaseType{"Binary", ascii.Binary, ascii.IsBinary},
	testCaseType{"Octal", ascii.Octal, ascii.IsOctal},
	testCaseType{"Hex", ascii.Hex, ascii.IsHex},
	testCaseType{"Latin", ascii.Latin, ascii.IsLatin},
	testCaseType{"Snake", ascii.Snake, ascii.IsSnake},
	testCaseType{"ASCII", ascii.ASCII, ascii.IsASCII},
	testCaseType{"Byte", ascii.Byte, ascii.IsByte},
}

func TestFilters(t *testing.T) {
	cases := make([]assert.F, len(testCases))
	for i, testCase := range testCases {
		cases[i] = assert.C(testCase.Name, filterTest(testCase.Set, testCase.Filter))
	}
	assert.Apply(t, cases...)
}
