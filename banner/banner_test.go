package banner

import (
	"testing"
)

func TestPrintHorizontalEdges(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"Hello, world!", "+-----------------*"},
		{"Lorem ipsum dolor sit amet", "+------------------------------*"},
	}

	b := NewDefaultBanner("")

	for _, c := range cases {
		b.Message = c.in
		if b.PrintHorizontalEdges() != c.out {
			t.Errorf("PrintTop(b) expected \"%s\". Got: \"%s\"", c.in, c.out)
		}
	}
}

func BenchmarkPrintHorizontalEdges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDefaultBanner("Hello, World").PrintHorizontalEdges()
	}
}

func TestPrintMiddle(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"Hello, world!", "|  Hello, world!  |"},
		{"Lorem ipsum dolor sit amet", "|  Lorem ipsum dolor sit amet  |"},
	}

	b := NewDefaultBanner("")

	for _, c := range cases {
		b.Message = c.in
		if b.PrintMiddle() != c.out {
			t.Errorf("PrintTop(b) expected \"%s\". Got: \"%s\"", c.in, c.out)
		}
	}
}

func BenchmarkPrintMiddle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDefaultBanner("Hello, World").PrintMiddle()
	}
}
