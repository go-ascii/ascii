package ascii

// Seq generates a sequence of consecutive bytes from begin to end (inclusive).
func Seq(begin, end byte) []byte {
	n := int(end-begin) + 1
	p := make([]byte, n)
	for i := range p {
		p[i] = begin + byte(i)
	}
	return p
}

// Sets of ASCII characters grouped by common traits.
var (
	Null    = []byte{0}
	Graphic = Seq(32, 126)
	Control = append(Seq(0, 31), 127)
	Space   = []byte{' ', '\t', '\n', '\v', '\f', '\r'}
	Quote   = []byte{'\'', '"', '`'}
	Upper   = Seq('A', 'Z')
	Lower   = Seq('a', 'z')
	Letter  = append(Upper, Lower...)
	Digit   = Seq('0', '9')
	NonZero = Seq('1', '9')
	Binary  = []byte{'0', '1'}
	Octal   = Seq('0', '7')
	Hex     = append(Digit, append(Seq('A', 'F'), Seq('a', 'f')...)...)
	Latin   = append(Letter, Digit...)
	Snake   = append(Latin, '_')
	ASCII   = Seq(0, 127)
	Byte    = Seq(0, 255)
)
