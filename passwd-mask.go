package passwdmask

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"strconv"
	"fmt"
)

var (
	alpha_lower = []byte("abcdefghijklmnopqrstuvwxyz")
	alpha_upper = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alpha_mixed = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	vowels_lower = []byte("aeiouy")
	vowels_upper = []byte("AEIOUY")
	vowels_mixed = []byte("aeiouyAEIOUY")
	consonants_lower = []byte("abcdfghjklmnpqrstvwxz")
	consonants_upper = []byte("ABCDFGHJKLMNPQRSTVWXZ")
	consonants_mixed = []byte("abcdfghjklmnpqrstvwxzABCDFGHJKLMNPQRSTVWXZ")

	numbers = []byte("0123456789")

	base64 = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/")

	hex_lower = []byte("abcdef0123456789")
	hex_upper = []byte("ABCDEF0123456789")

	alpha_numeric_lower = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	alpha_numeric_upper = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	alpha_numeric_mixed = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	number_symbols = []byte("!@#$%^&*()")
	dna = []byte("UCAG")

	code = map[rune][]byte{
		'a': alpha_lower,
		'A': alpha_upper,
		'B': alpha_mixed,
		'v': vowels_lower,
		'V': vowels_upper,
		'U': vowels_mixed,
		'c': consonants_lower,
		'C': consonants_upper,
		'D': consonants_mixed,
		'#': numbers,
		'n': alpha_numeric_lower,
		'N': alpha_numeric_upper,
		'M': alpha_numeric_mixed,
		'h': hex_lower,
		'H': hex_upper,
		'b': base64,
		's': number_symbols,
		'd': dna,
	}
)

func randomByteFrom(b []byte) byte {
	max := big.NewInt(int64(len(b)))
	index, _ := rand.Int(rand.Reader, max)
	return b[index.Uint64()]
}

func Generate (mask []byte) ([]byte, error) {
	var password bytes.Buffer
	var insert_char byte
	var slice, previous_set, custom_set, numbers []byte
	var x, count uint64
	var ok, inside_set_definition, inside_count_definition, previous_is_custom bool
	var parse_err error

	for i, char := range mask {
		switch char {

		case '[': // define set
			if inside_set_definition || inside_count_definition {
				return nil, fmt.Errorf("Set: Nested delimiter detected at character index %d.", i);
			}

			inside_set_definition = true
			custom_set = make([]byte, 0)

		case ']': // close set
			if !inside_set_definition || inside_count_definition {
				return nil, fmt.Errorf("Nested closing delimiter detected at character index %d.", i);
			}

			inside_set_definition = false
			previous_set = custom_set
			previous_is_custom = true

		case '{': // define expansion
			if inside_count_definition || inside_set_definition {
				return nil, fmt.Errorf("Count: Nested delimiter detected at character index %d.", i);
			}

			inside_count_definition = true
			numbers = make([]byte, 0)

		case '}': // close expansion
			if !inside_count_definition || inside_set_definition {
				return nil, fmt.Errorf("Nested closing delimiter detected at character index %d.", i);
			}

			inside_count_definition = false
			count, parse_err = strconv.ParseUint(string(numbers), 10, 64)

			if parse_err != nil {
				return nil, fmt.Errorf("'%s' could not be parsed as an integer.", numbers)
			}

			if !previous_is_custom {
				count = count - 1
			}

			for x = 0; x < count; x++ {
				password.WriteByte(randomByteFrom(previous_set))
			}

		default:
			if inside_set_definition {
				custom_set = append(custom_set, char)
			} else if inside_count_definition {
				numbers = append(numbers, char)
			} else {
				slice, ok = code[rune(char)]

				if !ok {
					insert_char = char
					previous_set = []byte{insert_char}
					previous_is_custom = false
					password.WriteByte(insert_char)
				} else {
					previous_set = slice
					previous_is_custom = false
					password.WriteByte(randomByteFrom(slice))
				}
			}
		}
	}

	return password.Bytes(), nil
}
