package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

var (
	mask             = flag.String("m", "h{16}", "Password generation mask. A string containing a variety of placeholders.")
	special_string   = flag.String("s", "!@#$%^&*_-.", "Special characters to use for 's' placeholder.")
	suppress_newline = flag.Bool("n", false, "If passed, suppress newline after outputting generated password.")

	r  = rand.New(rand.NewSource(time.Now().UnixNano()))
	re = regexp.MustCompile(`([aA#nNmhHbs]{1}){(\d+)}`)
)

func randomByteFrom(b []byte) byte {
	return b[r.Intn(len(b))]
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

func main() {
	flag.Usage = func() {
		fmt.Println("A password generator that uses masks.")
		flag.PrintDefaults()
		fmt.Println("The mask flag can consist of the following letters:")
		fmt.Println("  a: alpha, lowercase")
		fmt.Println("  A: alpha, uppercase")
		fmt.Println("  #: numbers")
		fmt.Println("  n: alpha-numeric, lowercase")
		fmt.Println("  N: alpha-numeric, uppercase")
		fmt.Println("  M: alpha-numeric, mixedcase")
		fmt.Println("  h: hex, lowercase")
		fmt.Println("  H: hex, uppercase")
		fmt.Println("  b: base64")
		fmt.Println("  s: specials")
		fmt.Println("All other characters will be untouched.")
		fmt.Println("{integer} repeats the previous character of the above list that many times, i.e. a{5} becomes aaaaa")
	}

	flag.Parse()

	alpha_lower := []byte("abcdefghijklmnopqrstuvwxyz")
	alpha_upper := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []byte("0123456789")
	specials := []byte(*special_string)

	base64 := []byte("+/")
	base64 = append(base64, alpha_lower...)
	base64 = append(base64, alpha_upper...)
	base64 = append(base64, numbers...)

	hex_lower := []byte("abcdef")
	hex_lower = append(hex_lower, numbers...)

	hex_upper := []byte("ABCDEF")
	hex_upper = append(hex_upper, numbers...)

	alpha_numeric_lower := make([]byte, 36)
	alpha_numeric_lower = append(alpha_numeric_lower, alpha_lower...)
	alpha_numeric_lower = append(alpha_numeric_lower, numbers...)

	alpha_numeric_upper := make([]byte, 36)
	alpha_numeric_upper = append(alpha_numeric_upper, alpha_upper...)
	alpha_numeric_upper = append(alpha_numeric_upper, numbers...)

	alpha_numeric_mixed := make([]byte, 62)
	alpha_numeric_mixed = append(alpha_numeric_mixed, alpha_upper...)
	alpha_numeric_mixed = append(alpha_numeric_mixed, alpha_lower...)
	alpha_numeric_mixed = append(alpha_numeric_mixed, numbers...)

	code := map[rune][]byte{
		'a': alpha_lower,
		'A': alpha_upper,
		'#': numbers,
		'n': alpha_numeric_lower,
		'N': alpha_numeric_upper,
		'M': alpha_numeric_mixed,
		'h': hex_lower,
		'H': hex_upper,
		'b': base64,
		's': specials,
	}

	var password bytes.Buffer
	var insert_char byte
	var slice []byte
	var ok bool

	expanded_mask := re.ReplaceAllFunc([]byte(*mask), replaceGroup)

	for _, char := range expanded_mask {
		slice, ok = code[rune(char)]

		if !ok {
			insert_char = char
		} else {
			insert_char = randomByteFrom(slice)
		}

		password.WriteByte(insert_char)
	}

	password.WriteTo(os.Stdout)

	if !*suppress_newline {
		fmt.Println()
	}
}
