package main

import (
	"flag"
	"fmt"
	"github.com/Lavos/passwd-mask"
)

var (
	mask             = flag.String("m", "h{16}", "Password generation mask. A string containing a variety of placeholders.")
	special_string   = flag.String("s", "!@#$%^&*_-.", "Special characters to use for 's' placeholder.")
	suppress_newline = flag.Bool("n", false, "If passed, suppress newline after outputting generated password.")
)

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

	password := passwdmask.Generate([]byte(*mask), []byte(*special_string))

	if *suppress_newline {
		fmt.Printf("%s", password)
	} else {
		fmt.Printf("%s\n", password)
	}
}
