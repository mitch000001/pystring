package pystring

import (
	"os/exec"
	"testing"
	"fmt"
	"strings"
)

func TestCapitalize(t *testing.T) {
	const in, out = "hello", "Hello"
	result := New(in).Capitalize().String()
	if out != result {
		t.Errorf("Capitalized version of %v did not become %v, but %v!\n", in, out, result)
	}
}

func TestStrip(t *testing.T) {
	const in, out = "    \n\n\nhello\t\t\n\n\n", "hello"
	result := New(in).Strip().String()
	if out != result {
		t.Errorf("Stripped version of %v did not become %v, but %v!\n", in, out, result)
	}
}

func TestCount(t *testing.T) {
	const in, out = "aabbaaccaaddaaffaa", 5
	result := New(in).Count("aa")
	if out != result {
		t.Errorf("Count of aa in %v, was not %v but %v!\n", in, out, result)
	}
}

func TestIndex(t *testing.T) {
	const in, out = "aabb", 2
	result, err := New(in).Index("bb")
	if (err != nil) || (out != result) {
		t.Errorf("Index of bb in %v, was not %v but %v!\n", in, out, result)
	}
}

func TestEndsWith(t *testing.T) {
	const in1, in2 = "aabbcc", "abcdef"
	if New(in1).EndsWith("cc") == false {
		t.Errorf("EndsWith: %v does end with cc!\n", in1)
	}
	if New(in2).EndsWith("dd") == true {
		t.Errorf("EndsWith: %v does not end with dd!\n", in2)
	}
}

func TestStartsWith(t *testing.T) {
	const in1, in2 = "aabbcc", "abcdef"
	if New(in1).StartsWith("cc") != false {
		t.Errorf("StartsWith: %v does not start with cc!\n", in1)
	}
	if New(in2).StartsWith("ab") != true {
		t.Errorf("StartsWith: %v does start with ab!\n", in2)
	}
}

func TestFind(t *testing.T) {
	const in = "aabbcc"
	if New(in).Find("cc") != 4 {
		t.Errorf("Find: %v has cc in position 4!\n", in)
	}
	if New(in).Find("cd") != -1 {
		t.Errorf("Find: %v does not have cd!\n", in)
	}
}

func TestIsDigit(t *testing.T) {
	const in1, in2 = "223098473", "x1z"
	if New(in1).IsDigit() != true {
		t.Errorf("IsDigit: %v should be true!\n", in1)
	}
	if New(in2).IsDigit() != false {
		t.Errorf("IsDigit: %v should be false!\n", in2)
	}
	if New("").IsDigit() != false {
		t.Errorf("IsDigit: empty string should be false!\n")
	}
}

func TestAdd(t *testing.T) {
	a := New("abc")
	if a.Add("cde").String() != "abccde" {
		t.Errorf("abc + cde should be abccde\n")
	}
}

func TestSubstract(t *testing.T) {
	a := New("abcdef")
	if a.Subtract("def") != "abc" {
		t.Errorf("abcdef - def should be abc\n")
	}
	a = New("ost kake ost kake ost")
	if a.Subtract("kake") != "ost kake ost  ost" {
		t.Errorf("ost kake ost kake ost - kake should be ost kake ost  ost\n")
	}
}

func TestSplit(t *testing.T) {
	a := New("a:b:c")
	if a.Split(":")[1] != "b" {
		t.Errorf("a:b:c split on : should be a b c\n")
	}
}

func TestMultiply(t *testing.T) {
	a := New("a")
	if a.Multiply(5) != "aaaaa" {
		t.Errorf("a*5 should be aaaaa\n")
	}
	b := New("ost")
	if b.Multiply(2) != "ostost" {
		t.Errorf("ost*2 should be ostost\n")
	}
}

func TestMultiplyUnicode(t *testing.T) {
	a := New("┐")
	if a.Multiply(5) != "┐┐┐┐┐" {
		t.Errorf("a*5 should be ┐┐┐┐┐\n")
	}
}

/* Checks if the constants are equal to the ones in Python by running the python interpreter.
   Also exercises the New(), Join() and Encode() functions.
*/
func TestConstants(t *testing.T) {
	var shouldBeBytes []byte
	var cmd *exec.Cmd
	constants := []string{"ascii_letters", "ascii_lowercase", "ascii_uppercase", "digits", "hexdigits", "octdigits", "punctuation", "printable", "whitespace"}
	shouldbe := []string{ASCII_letters, ASCII_lowercase, ASCII_uppercase, Digits, HexDigits, OctDigits, Punctuation, Printable, Whitespace}
	for i, constant := range constants {
		cmd = exec.Command("python", "-c", "import string; print(string."+constant+")")
		output, err := cmd.Output()
		if err != nil {
			/* One of the commands failed, assume Python is not available */
			// This can be used if one doesn't want to ignore the lack of python:
			//t.Errorf("execution failed: %s %s\n", New(" ").Join(cmd.Args), output)
			return
		}
		shouldBeBytes = New(shouldbe[i]).Encode()
		for i2, b := range shouldBeBytes {
			if (i2 >= len(output)) || (b != output[i2]) {
				if len(output) == 0 {
					output = New("no output").Encode()
				}
				t.Errorf("constant %s failed, got: %s\n", constants[i], output)
				return
			}
		}
	}
}

func TestReplace(t *testing.T) {
	a := New("hello")
	a.Replace("l", "p")
	if a.String() != "heppo" {
		t.Errorf("hello with l replaced should be heppo\n")
	}
}

func TestSplitLines(t *testing.T) {
	s := "First line, with LF\nSecond line, with CR\rThird line, with CRLF\r\n" +
	     "Two blank lines with LFs\n\n\nTwo blank lines with CRs\r\r\rTwo blank" +
		 "lines with CRLFs\r\n\r\n\r\nThree blank lines with a jumble of things:" +
		 "\r\n\r\r\n\nEnd without a newline."
	s2 := "First line, with LF\nSecond line, with CR\nThird line, with CRLF\n" +
	     "Two blank lines with LFs\nTwo blank lines with CRs\nTwo blank" +
		 "lines with CRLFs\nThree blank lines with a jumble of things:" +
		 "\nEnd without a newline."
	a := New(s)
	lines := a.SplitLines()
	joined := strings.Join(lines, "\n")
	fmt.Println(joined)
	if joined != s2 {
		t.Errorf("SplitLines does not work correctly, is %s, but should be %s.\n", joined, s2)
	}
}
