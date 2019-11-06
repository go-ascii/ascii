package ascii_test

import (
	"testing"

	"github.com/ktnyt/ascii"
	"github.com/stretchr/testify/require"
)

func filterTest(set []byte, filter ascii.Filter) func(*testing.T) {
	return func(t *testing.T) {
		for b := byte(0); b < 128; b++ {
			if ascii.Contains(set)(b) {
				require.True(t, filter(b))
			}
		}
	}
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
}

func TestFilters(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.Name, filterTest(testCase.Set, testCase.Filter))
	}
}
