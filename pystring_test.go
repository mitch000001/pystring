package pystring

import "testing"

func TestCapitalize(t *testing.T) {
	const in, out = "hello", "Hello"
	result := New(in).Capitalize().s
	if out != result {
		t.Errorf("Capitalized version of %v did not become %v, but %v!\n", in, out, result)
	}
}

func TestStrip(t *testing.T) {
	const in, out = "    \n\n\nhello\t\t\n\n\n", "hello"
	result := New(in).Strip().s
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
