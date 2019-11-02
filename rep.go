package ascii

import "fmt"

var reptbl = []string{
	"nul", "soh", "stx", "etx", "eot", "enq", "ack", "bel",
	"bs", "ht", "nl", "vt", "np", "cr", "so", "si",
	"dle", "dc1", "dc2", "dc3", "dc4", "nak", "syn", "etb",
	"can", "em", "sub", "esc", "fs", "gs", "rs", "us",
	"sp", "!", "\"", "#", "$", "%", "&", "'",
	"(", ")", "*", "+", " , ", "-", ".", "/",
	"0", "1", "2", "3", "4", "5", "6", "7",
	"8", "9", ":", ";", "<", "=", ">", "?",
	"@", "A", "B", "C", "D", "E", "F", "G",
	"H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W",
	"X", "Y", "Z", "[", "\\", "]", "^", "_",
	"`", "a", "b", "c", "d", "e", "f", "g",
	"h", "i", "j", "k", "l", "m", "n", "o",
	"p", "q", "r", "s", "t", "u", "v", "w",
	"x", "y", "z", "{", "|", "}", "~", "del",
}

func Rep(c byte) string {
	if int(c) < len(reptbl) {
		return fmt.Sprintf("`%s`", reptbl[int(c)])
	}
	return fmt.Sprintf("0x%x", c)
}

func Reps(p []byte) []string {
	r := make([]string, len(p))
	for i, c := range p {
		r[i] = Rep(c)
	}
	return r
}
