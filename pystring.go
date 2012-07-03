/*
 * pystring
 * Python-like strings for Go
 * Alexander Rødseth <rodseth@gmail.com>
 * july 2012
 * GPL2
 */

package pystring

import (
	"errors"
	"strings"
)

const (
	ascii_letters   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ascii_lowercase = "abcdefghijklmnopqrstuvwxyz"
	ascii_uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits          = "0123456789"
	hexdigits       = "0123456789abcdefABCDEF"
	octdigits       = "01234567"
	punctuation     = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	printable       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c"
	whitespace      = " \t\n\r\x0b\x0c"
)

type PyString struct {
	s string
}

func New(s string) *PyString {
	return &PyString{s}
}

func (p *PyString) Capitalize() *PyString {
	p.s = strings.Title(p.s)
	return p
}

func (p *PyString) Strip() *PyString {
	p.s = strings.TrimSpace(p.s)
	return p
}

func (p *PyString) Count(s string) int {
	return strings.Count(p.s, s)
}

func (p *PyString) Encode() []byte {
	return []byte(p.s)
}

func (p *PyString) Index(s string) (int, error) {
	var err error
	i := strings.Index(p.s, s)
	if i == -1 {
		err = errors.New("Not found")
	} else {
		err = nil
	}
	return i, err
}
