package ascii

import "strings"

// Filter represents a byte filter fucntion type.
type Filter func(byte) bool

// AsFilter creates a filter that will match only the given byte.
func AsFilter(b byte) Filter { return func(c byte) bool { return b == c } }

// Range creates a filter that will match any byte in between (inclusive).
func Range(begin, end byte) Filter {
	return func(c byte) bool { return begin <= c && c <= end }
}

// Contains creates a filter that will match one of the given bytes.
func Contains(p []byte) Filter {
	s := string(p)
	return func(c byte) bool { return strings.IndexByte(s, c) >= 0 }
}

// Or will match if any one of the given filters match.
func Or(filters ...Filter) Filter {
	return func(c byte) bool {
		for _, filter := range filters {
			if filter(c) {
				return true
			}
		}
		return false
	}
}

// Filters for ASCII character sets.
var (
	IsNullFilter    = AsFilter(0)
	IsControlFilter = Range(0, 31)
	IsSpaceFilter   = Contains(Space)
	IsUpperFilter   = Range('A', 'Z')
	IsLowerFilter   = Range('a', 'z')
	IsLetterFilter  = Or(IsUpperFilter, IsLowerFilter)
	IsDigitFilter   = Range('0', '9')
	IsNonZeroFilter = Range('1', '9')
	IsLatinFilter   = Or(IsLetterFilter, IsDigitFilter)
	IsSnakeFilter   = Or(IsLatinFilter, AsFilter('_'))
	IsASCIIFilter   = Range(0, 127)
)

// IsNull matches a null character.
func IsNull(c byte) bool { return IsNullFilter(c) }

// IsControl matches a control byte.
func IsControl(c byte) bool { return IsControlFilter(c) }

// IsSpace matches a whitespace character.
func IsSpace(c byte) bool { return IsSpaceFilter(c) }

// IsUpper matches an uppercase letter.
func IsUpper(c byte) bool { return IsUpperFilter(c) }

// IsLower matches a lowercase letter.
func IsLower(c byte) bool { return IsLowerFilter(c) }

// IsLetter matches an uppercase or lowercase letter.
func IsLetter(c byte) bool { return IsLetterFilter(c) }

// IsDigit matches a digit character.
func IsDigit(c byte) bool { return IsDigitFilter(c) }

// IsNonZero matches a non-zero digit character.
func IsNonZero(c byte) bool { return IsNonZeroFilter(c) }

// IsLatin matches a letter or a digit character.
func IsLatin(c byte) bool { return IsLatinFilter(c) }

// IsSnake matches a latin character or an underscore.
func IsSnake(c byte) bool { return IsSnakeFilter(c) }

// IsASCII matcehs any ASCII character.
func IsASCII(c byte) bool { return IsASCIIFilter(c) }
