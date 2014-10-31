A small password generator that uses masks. It's just using `crypto/rand` internally, so I should be *pretty* random. I wouldn't protect government secrets with this at any rate.

## Building

```bash
go get github.com/Lavos/passwd-mask
cd $GOPATH/src/github.com/Lavos/passwd-mask/commands
go build .
```

## Usage
```bash
passwd-mask mask [-n]
```

### mask
The mask value is a string with placeholder characters that represent a larger set of characters. Each instance is replaced with a random character from the chosen set.

The characters for each set are:

* `a`: alpha, lowercase
* `A`: alpha, uppercase
* `#`: numbers
* `n`: alpha-numeric, lowercase 
* `N`: alpha-numeric, uppercase
* `M`: alpha-numeric, mixed-case
* `h`: hex, lowercase
* `H`: hex, uppercase
* `b`: base64

Characters not in this set are copied to the generated password untouched.

{integer} can be used to repeat previous character `integer` times. For example, "h{5}" expands to "hhhhh".
[characters]{integer} defines a custom set of characters to be used integer times, i.e. [abc]{5} becomes cbaba"

Passing `-` as the mask value makes the program read from STDIN.

#### Examples
```bash
passwd-mask "aaa_AAA_###"
returns: ozj_DNQ_646
```
3x lowercase alpha _ 3x uppercase alpha _ 3x numbers

```bash
passwd-mask "bbbbb-#####"
returns: +KKRK-67255
```
5x base64 - 5x numbers

```bash
passwd-mask "b{20}"
returns: We//QplIEu50bsilLkLR9whwy
```
20 characters of base64.

```bash
passwd-mask "[.-]{5}"
returns: .--.-
```
5 characters from custom set of .-

### -n
if you using this command for scripting, you can suppress printing the newline at the end of the output by passing the `-n` flag.
