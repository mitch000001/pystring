/*
 * pystring
 * Python-like strings for Go
 * Alexander Rødseth <rodseth@gmail.com>
 * july 2012
 * GPL2
 */

package pystring

import (
	"bytes"
	"errors"
	"strings"
)

const (
	ASCII_letters   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ASCII_lowercase = "abcdefghijklmnopqrstuvwxyz"
	ASCII_uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits          = "0123456789"
	HexDigits       = "0123456789abcdefABCDEF"
	OctDigits       = "01234567"
	Punctuation     = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	Printable       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c"
	Whitespace      = " \t\n\r\x0b\x0c"
)

type PyString struct {
	text string
}

func New(text string) *PyString {
	return &PyString{text}
}

func (p *PyString) Capitalize() *PyString {
	p.text = strings.Title(p.text)
	return p
}

func (p *PyString) Strip() *PyString {
	p.text = strings.TrimSpace(p.text)
	return p
}

func (p *PyString) Count(text string) int {
	return strings.Count(p.text, text)
}

func (p *PyString) Encode() []byte {
	return []byte(p.text)
}

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

/* strings.HasSuffix should also work */
func (p *PyString) EndsWith(text string) bool {
	startpos := len(p.text) - len(text)
	if startpos < 0 {
		return false
	}
	for pos := startpos; pos < len(p.text); pos++ {
		if p.text[pos] != text[pos-startpos] {
			return false
		}
	}
	return true
}

/* strings.HasPrefix should also work */
func (p *PyString) StartsWith(text string) bool {
	if len(text) > len(p.text) {
		return false
	}
	for pos := 0; pos < len(text); pos++ {
		if p.text[pos] != text[pos] {
			return false
		}
	}
	return true
}

/* Instead of "xyz" in x in Python, use x.Has("xyz") in Go */
func (p *PyString) Has(text string) bool {
	return -1 != strings.Index(p.text, text)
}

/* To fetch the string */
func (p *PyString) Get() string {
	return p.text
}

/* Is the string empty? Instead of if s: ... in Python, use if s.Empty() { ... in Go */
func (p *PyString) Empty() bool {
	return 0 == len(p.text)
}

func (p *PyString) Find(text string) int {
	return strings.Index(p.text, text)
}

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

/* Instead of a + b in Python, use a.Add(b) in Go */
func (p *PyString) Add(text string) string {
	p.text += text
	return p.text
}

/* Append a string */
func (p *PyString) Append(s string) {
	p.text += s
}

/* Remove the last occurence of a string */
func (p *PyString) Subtract(text string) string {
	pos := p.RFind(text)
	if pos != -1 {
		p.text = p.text[:pos] + p.text[pos+len(text):]
	}
	return p.text
}

/* Find the last occurance */
func (p *PyString) RFind(text string) int {
	return strings.LastIndex(p.text, text)
}

func (p *PyString) Split(sep string) []string {
	return strings.Split(p.text, sep)
}

func (p *PyString) Multiply(n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(p.text)
	}
	return buf.String()
}
