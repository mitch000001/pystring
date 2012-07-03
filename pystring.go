/*
 * pystring
 * Python-like strings for Go
 * Alexander Rødseth <rodseth@gmail.com>
 * july 2012
 * GPL2
 */

package pystring

type PyString struct {
	s string
}

func (p *PyString) strip() *PyString {
	return p
}

func (p *PyString) empty() bool {
	return "" == p.s
}
