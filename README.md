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
passwd-mask -m "b{2048}"
returns: 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
```

2048 characters of base64.

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

