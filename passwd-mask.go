package main

import (
	"os"
	"flag"
	"math/rand"
	"bytes"
	"time"
	"fmt"
)

var (
	mask = flag.String("m", "hhhhhhhhhhhh", "Password generation mask.")
	special_string = flag.String("s", "!@#$%^&*_-.", "User specified special characters.")

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func randomByteFrom(b []byte) byte {
	return b[r.Intn(len(b))]
}

func main () {
	flag.Usage = func () {
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
	}

	flag.Parse()

	alpha_lower := []byte("abcdefghijklmnopqrstuvwxyz")
	alpha_upper := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []byte("0123456789")
	base64_extras := []byte("+/")
	specials := []byte(*special_string)

	base64 := append(alpha_lower, alpha_upper...)
	base64 = append(base64, numbers...)
	base64 = append(base64, base64_extras...)

	hex_lower := append(numbers, []byte("abcef")...)
	hex_upper := append(numbers, []byte("ABCDEF")...)

	alpha_numeric_lower := append(alpha_lower, numbers...)
	alpha_numeric_upper := append(alpha_upper, numbers...)
	alpha_numeric_mixed := append(alpha_lower,  alpha_upper...)
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

	for _, char := range *mask {
		slice, ok = code[char]

		if !ok {
			insert_char = byte(char)
		} else {
			insert_char = randomByteFrom(slice)
		}

		password.WriteByte(insert_char)
	}

	password.WriteTo(os.Stdout)
}
