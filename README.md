A small password generator that uses masks. It's just using `crypto/rand` internally, so I should be *pretty* random. I wouldn't protect government secrets with this at any rate.

## Building

```bash
go get github.com/Lavos/passwd-mask
cd $GOPATH/src/github.com/Lavos/passwd-mask/commands
go build .
```

## Usage
```bash
passwd-mask [-m mask] [-s specials] [-n]
```

### -m mask
In most usage cases, you will want to pass your specific mask. Otherwise, `passwd-mask` defaults to "h{16}".

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
* `s`: specials

Characters not in this set are copied to the generated password untouched.

{integer} can be used to repeat previous character `integer` times. For example, "h{5}" expands to "hhhhh".

#### Examples
```bash
passwd-mask -m "aaa_AAA_###"
returns: ozj_DNQ_646
```
3x lowercase alpha _ 3x uppercase alpha _ 3x numbers

```bash
passwd-mask -m "bbbbb-#####"
returns: +KKRK-67255
```
5x base64 - 5x numbers

```bash
passwd-mask -m "b{20}"
returns: We//QplIEu50bsilLkLR9whwy
```
20 characters of base64.

### -s specials

You can specify your own special characters for the case that your provider only accepts specific characters.

#### Examples

```bash
passwd-mask -m "aaaaa-sssss" -s "#$"
returns: spgcr-$##$#
```
5x lowercase alpha - 5x of either # or $

### -n

If you using this command for scripting, you can suppress printing the newline at the end of the output by passing the `-n` flag.
