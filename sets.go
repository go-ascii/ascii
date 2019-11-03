package ascii

// Seq generates a sequence of consecutive bytes from begin to end (inclusive).
func Seq(begin, end byte) []byte {
	n := end - begin + 1
	p := make([]byte, n)
	for i := byte(0); i < n; i++ {
		p[i] = begin + i
	}
	return p
}

// Sets of ASCII characters grouped by common traits.
var (
	Null    = []byte{0}
	Control = Seq(0, 31)
	Space   = []byte{' ', '\t', '\n', '\v', '\f', '\r'}
	Upper   = Seq('A', 'Z')
	Lower   = Seq('a', 'z')
	Letter  = append(Upper, Lower...)
	Digit   = Seq('0', '9')
	NonZero = Seq('1', '9')
	Latin   = append(Letter, Digit...)
	Snake   = append(Latin, '_')
	ASCII   = Seq(0, 127)
)
