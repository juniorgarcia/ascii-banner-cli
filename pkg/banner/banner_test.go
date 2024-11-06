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
		r, err := b.PrintHorizontalEdges(HorLineTypeTop)
		if err != nil {
			t.Errorf("PrintTop(b) generated error: %s", err)
		}
		if r != c.out {
			t.Errorf("PrintTop(b) expected \"%s\". Got: \"%s\"", c.in, c.out)
		}
	}
}

func BenchmarkPrintHorizontalEdges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewDefaultBanner("Hello, World").PrintHorizontalEdges(HorLineTypeBottom)
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
		NewDefaultBanner("Hello, world!").PrintMiddle()
	}
}

func TestPrintFullBanner(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{
			"Hello, world!",
			"+-----------------*\n" +
				"|  Hello, world!  |\n" +
				"*-----------------+",
		},
		{
			"Foo Bar",
			"+-----------*\n" +
				"|  Foo Bar  |\n" +
				"*-----------+",
		},
	}

	b := NewDefaultBanner("")

	for _, c := range cases {
		b.Message = c.in
		banner, err := b.PrintFullBanner()
		if err != nil {
			t.Errorf("PrintFullBanner() raised an error: %+v", err)
		}
		if banner != c.out {
			t.Errorf("PrintFullBanner() expected: \n\"%s\". \n Got: \n\"%s\"", c.out, banner)
		}
	}
}

func BenchmarkPrintFullBanner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewDefaultBanner("Hello, world!").PrintFullBanner()
	}
}
