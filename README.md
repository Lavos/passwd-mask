A small password generator that uses masks. It's just using Rand.int internally, so I wouldn't use it for anything that needs **truely** cryptographically random digits.

## Building

```bash
go get github.com/Lavos/passwd-mask
cd $GOPATH/src/github.com/Lavos/passwd-mask
go build .
```

## Usage

```bash
passwd-mask [-m mask] [-s specials]
```

### -m mask
In most usage cases, you will want to pass your specific mask. Otherwise, `passwd-mask` defaults to "hhhhhhhhhh".

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

### -s specials

You can specify your own special characters for the case that your provider only accepts specific characters.

#### Examples

```bash
passwd-mask -m "aaaaa-sssss" -s "#$"
returns: spgcr-$##$#
```
5x lowercase alpha - 5x of either # or $

```bash
passwd-mask -m "ssssssssssssss" -s ".-"
returns: --..-.-.-..-.-
```

15x dots or slashes

