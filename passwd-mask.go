package passwdmask

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"regexp"
	"strconv"
)

var (
	re = regexp.MustCompile(`([aA#nNmhHbs]{1}){(\d+)}`)
	alpha_lower = []byte("abcdefghijklmnopqrstuvwxyz")
	alpha_upper = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers = []byte("0123456789")
	base64 = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/")
	hex_lower = []byte("abcdef0123456789")
	hex_upper = []byte("ABCDEF0123456789")
	alpha_numeric_lower = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	alpha_numeric_upper = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	alpha_numeric_mixed = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	code = map[rune][]byte{
		'a': alpha_lower,
		'A': alpha_upper,
		'#': numbers,
		'n': alpha_numeric_lower,
		'N': alpha_numeric_upper,
		'M': alpha_numeric_mixed,
		'h': hex_lower,
		'H': hex_upper,
		'b': base64,
	}
)

func randomByteFrom(b []byte) byte {
	max := big.NewInt(int64(len(b)))
	index, _ := rand.Int(rand.Reader, max)
	return b[index.Uint64()]
}

func replaceGroup(b []byte) []byte {
	matches := re.FindSubmatch(b)

	if matches == nil {
		// no matches found shouldn't happen because it wouldn't get
		// through the regexp
		return b
	}

	char := matches[1]
	count, err := strconv.ParseInt(string(matches[2]), 10, 64)

	if err != nil {
		// could not parse int, but it shouldn't happen because
		// then it wouldn't get through the regexp
		return b
	}

	var buf bytes.Buffer
	var x int64

	for ; x < count; x++ {
		buf.Write(char)
	}

	return buf.Bytes()
}

func Generate (mask, specials []byte) []byte {
	var password bytes.Buffer
	var insert_char byte
	var slice []byte
	var ok bool

	expanded_mask := re.ReplaceAllFunc(mask, replaceGroup)

	for _, char := range expanded_mask {
		if char == 's' {
			insert_char = randomByteFrom(specials)
		} else {
			slice, ok = code[rune(char)]

			if !ok {
				insert_char = char
			} else {
				insert_char = randomByteFrom(slice)
			}
		}

		password.WriteByte(insert_char)
	}

	return password.Bytes()
}
