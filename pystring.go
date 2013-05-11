package pystring

/*
 * pystring
 * Python-like string methods for Go
 * Alexander Rødseth <rodseth@gmail.com>
 * July 2012
 * GPL2
 */

import (
	"bytes"
	"errors"
	"strings"
)

const (
	/* string.ascii_letters */
	ASCII_letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	/* string.ascii_lowercase */
	ASCII_lowercase = "abcdefghijklmnopqrstuvwxyz"
	/* string.ascii_uppercase */
	ASCII_uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	/* string.digits */
	Digits = "0123456789"
	/* string.hexdigits */
	HexDigits = "0123456789abcdefABCDEF"
	/* string.octdigits */
	OctDigits = "01234567"
	/* string.punctuation */
	Punctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	/* string.printable */
	Printable = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c"
	/* string.whitespace */
	Whitespace = " \t\n\r\x0b\x0c"
)

type PyString struct {
	text string
}

/*
 * Helper functions, specific to the PyString type
 */

func New(text string) *PyString {
	return &PyString{text}
}

/* To fetch the string */
func (p *PyString) Get() string {
	return p.text
}

/*
 * String methods, in alphabetical order
 * These are aproximately the same as in Python, some are simpler
 */

/* .capitalize() */
func (p *PyString) Capitalize() *PyString {
	p.text = strings.Title(p.text)
	return p
}

/* TODO .center() */
func (p *PyString) Center() *PyString {
	return p
}

/* .count() */
func (p *PyString) Count(text string) int {
	return strings.Count(p.text, text)
}

/* .encode() */
func (p *PyString) Encode() []byte {
	return []byte(p.text)
}

/* .endswith() */
func (p *PyString) EndsWith(text string) bool {
	return strings.HasSuffix(p.text, text)
}

/* TODO .expandtabs() */
func (p *PyString) ExpandTabs() *PyString {
	return p
}

/* .find() */
func (p *PyString) Find(text string) int {
	return strings.Index(p.text, text)
}

/* TODO .format() */
func (p *PyString) Format() *PyString {
	return p
}

/* TODO .format_map() */
func (p *PyString) FormatMap() *PyString {
	return p
}

/* .index() */
func (p *PyString) Index(text string) (int, error) {
	var err error
	i := strings.Index(p.text, text)
	if i == -1 {
		err = errors.New("Not found")
	} else {
		err = nil
	}
	return i, err
}

/* TODO .isalnum() */
func (p *PyString) IsAlNum() *PyString {
	return p
}

/* TODO .isalpha() */
func (p *PyString) IsAlpha() *PyString {
	return p
}

/* TODO .isdecimal() */
func (p *PyString) IsDecimal() *PyString {
	return p
}

/* .isdigit() */
func (p *PyString) IsDigit() bool {
	var isDigit bool
	if p.Empty() {
		return false
	}
	for _, letter := range p.text {
		isDigit = false
		for _, digit := range Digits {
			if letter == digit {
				isDigit = true
				break
			}
		}
		if !isDigit {
			return false
		}
	}
	return true
}

/* TODO .isidentifier() */
func (p *PyString) IsIdentifier(text string) bool {
	return false
}

/* TODO .islower() */
func (p *PyString) IsLower(text string) bool {
	return false
}

/* TODO .isnumeric() */
func (p *PyString) IsNumeric(text string) bool {
	return false
}

/* TODO .isprintable() */
func (p *PyString) IsPrintable(text string) bool {
	return false
}

/* TODO .isspace() */
func (p *PyString) IsSpace(text string) bool {
	return false
}

/* TODO .istitle() */
func (p *PyString) IsTitle(text string) bool {
	return false
}

/* TODO .isupper() */
func (p *PyString) IsUpper(text string) bool {
	return false
}

/* .join(), Join a list of strings, sep.join(sl) in Python */
func (p *PyString) Join(sl []string) string {
	return strings.Join(sl, p.text)
}

// TODO ljust, lower, lstrip, maketrans, partition, replace

/* .rfind() */
func (p *PyString) RFind(text string) int {
	return strings.LastIndex(p.text, text)
}

// TODO rindex, rjust, rpartition, rsplit, rstrip

/* .split() */
func (p *PyString) Split(sep string) []string {
	return strings.Split(p.text, sep)
}

// TODO .splitlines()

/* .startswith() */
func (p *PyString) StartsWith(text string) bool {
	return strings.HasPrefix(p.text, text)
}

/* .strip() */
func (p *PyString) Strip() *PyString {
	p.text = strings.TrimSpace(p.text)
	return p
}

// TODO: swapcase, title, translate, upper, zfill

/*
 * String functions that exists in Python as part of
 * the syntax, like Add() instead of "+"
 */

/* Instead of "a in b" in Python, use b.Has(a) */
func (p *PyString) Has(text string) bool {
	return -1 != strings.Index(p.text, text)
}

/* Instead of "a in b" in Python, there is also a.In(b) */
func (p *PyString) In(text string) bool {
	return -1 != strings.Index(text, p.text)
}

/* Instead of if a: ... in Python, use if a.Empty() */
func (p *PyString) Empty() bool {
	return "" == p.text
}

/* Instead of a + b in Python, use a.Add(b) */
func (p *PyString) Add(text string) *PyString {
	return New(p.text + text)
}

/* Instead of a += b in Python, use a.Append(b) */
func (p *PyString) Append(text string) {
	p.text += text
}

/* Instead of * in Python, use a.Multiply(n) */
func (p *PyString) Multiply(n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(p.text)
	}
	return buf.String()
}

/*
 * Functions with no direct equivivalent in Python
 */

/* Remove the last occurence of a string */
func (p *PyString) Subtract(text string) string {
	pos := p.RFind(text)
	if pos != -1 {
		return p.text[:pos] + p.text[pos+len(text):]
	}
	return p.text
}
