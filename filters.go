package ascii

import "strings"

// Filter represents a byte filter fucntion type.
type Filter func(byte) bool

// Range creates a filter that will match any byte in between (inclusive).
func Range(begin, end byte) Filter {
	return func(c byte) bool { return begin <= c && c <= end }
}

// Is creates a filter that will match only the given byte.
func Is(b byte) Filter { return func(c byte) bool { return b == c } }

// Contains creates a filter that will match one of the given bytes.
func Contains(p []byte) Filter {
	s := string(p)
	return func(c byte) bool { return strings.IndexByte(s, c) >= 0 }
}

// AsFilter attempts to create an optimized filter for the given arguments.
func AsFilter(args ...byte) Filter {
	switch len(args) {
	case 0:
		return Filter(func(byte) bool { return true })
	case 1:
		return Is(args[0])
	default:
		return Contains(args)
	}
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

// Not will invert the given filter.
func Not(filter Filter) Filter {
	return func(c byte) bool { return !filter(c) }
}

// Filters for ASCII character sets.
var (
	IsNullFilter    = AsFilter(0)
	IsGraphicFilter = Range(32, 126)
	IsControlFilter = Not(IsGraphicFilter)
	IsSpaceFilter   = Contains(Space)
	IsUpperFilter   = Range('A', 'Z')
	IsLowerFilter   = Range('a', 'z')
	IsLetterFilter  = Or(IsUpperFilter, IsLowerFilter)
	IsDigitFilter   = Range('0', '9')
	IsNonZeroFilter = Range('1', '9')
	IsBinaryFilter  = Or(AsFilter('0'), AsFilter('1'))
	IsOctalFilter   = Range('0', '7')
	IsHexFilter     = Or(IsDigitFilter, Range('A', 'F'), Range('a', 'f'))
	IsLatinFilter   = Or(IsLetterFilter, IsDigitFilter)
	IsSnakeFilter   = Or(IsLatinFilter, AsFilter('_'))
	IsASCIIFilter   = Range(0, 127)
	IsByteFilter    = Filter(func(byte) bool { return true })
)

// IsNull matches a null byte.
func IsNull(c byte) bool { return IsNullFilter(c) }

// IsGraphic matches a graphic character.
func IsGraphic(c byte) bool { return IsGraphicFilter(c) }

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

// IsBinary matches a binary digit character.
func IsBinary(c byte) bool { return IsBinaryFilter(c) }

// IsOctal matches a octal digit character.
func IsOctal(c byte) bool { return IsOctalFilter(c) }

// IsHex matches a hexadecimal digit character.
func IsHex(c byte) bool { return IsHexFilter(c) }

// IsNonZero matches a non-zero digit character.
func IsNonZero(c byte) bool { return IsNonZeroFilter(c) }

// IsLatin matches a letter or a digit character.
func IsLatin(c byte) bool { return IsLatinFilter(c) }

// IsSnake matches a latin character or an underscore.
func IsSnake(c byte) bool { return IsSnakeFilter(c) }

// IsASCII matches any ASCII character.
func IsASCII(c byte) bool { return IsASCIIFilter(c) }

// IsByte matches any Byte.
func IsByte(c byte) bool { return IsByteFilter(c) }
