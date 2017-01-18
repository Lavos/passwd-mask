package main

import (
	"flag"
	"fmt"
	"github.com/Lavos/passwd-mask"
	"os"
	"log"
	"bufio"
)

var (
	suppress_newline = flag.Bool("n", false, "If passed, suppress newline after outputting generated password.")
)

func main() {
	flag.Usage = func() {
		fmt.Println("A password generator that uses masks to meet specific requirements.")
		fmt.Println("Usage: passwd-mask mask")
		fmt.Println("If you want to read from STDIN, pass '-' as the mask value.")
		flag.PrintDefaults()
		fmt.Println("The mask can consist of the following letters:")
		fmt.Println("  a: alpha, lowercase")
		fmt.Println("  A: alpha, uppercase")
		fmt.Println("  B: alpha, mixedcase")
		fmt.Println("  v: vowels, lowercase")
		fmt.Println("  V: vowels, uppercase")
		fmt.Println("  U: vowels, mixedcase")
		fmt.Println("  c: consonants, lowercase")
		fmt.Println("  C: consonants, uppercase")
		fmt.Println("  D: consonants, mixedcase")
		fmt.Println("  #: numbers")
		fmt.Println("  n: alpha-numeric, lowercase")
		fmt.Println("  N: alpha-numeric, uppercase")
		fmt.Println("  M: alpha-numeric, mixedcase")
		fmt.Println("  h: hex, lowercase")
		fmt.Println("  H: hex, uppercase")
		fmt.Println("  b: base64")
		fmt.Println("  s: number symbols, ie: !@#$%^&*()")
		fmt.Println("  d: DNA sequence characters, ie: UCAG")
		fmt.Println("{integer} repeats the previous character of the above list that many times, i.e. a{5} becomes aaaaa")
		fmt.Println("[characters]{integer} defines a custom set of characters to be used integer times, i.e. [abc]{5} becomes cbaba")
		fmt.Println("All other characters will be untouched.")
	}

	flag.Parse()

	mask := flag.Arg(0)
	var password []byte
	var perr error

	if mask == "" {
		log.Fatal("No message supplied.")
	}

	if mask == "-" { // read from STDIN
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			password, perr = passwdmask.Generate(scanner.Bytes())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		password, perr = passwdmask.Generate([]byte(mask))
	}

	if perr != nil {
		log.Fatal(perr)
	}

	if *suppress_newline {
		fmt.Printf("%s", password)
	} else {
		fmt.Printf("%s\n", password)
	}
}
