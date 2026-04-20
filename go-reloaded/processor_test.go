package main

import "testing"

func TestHex(t *testing.T) {
	got := Processor("1E (hex) files were added")
	want := "30 files were added"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestBin(t *testing.T) {
	got := Processor("It has been 10 (bin) years")
	want := "It has been 2 years"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestUpper(t *testing.T) {
	got := Processor("go (up)")
	want := "GO"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestLower(t *testing.T) {
	got := Processor("STOP (low)")
	want := "stop"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCap(t *testing.T) {
	got := Processor("hello (cap)")
	want := "Hello"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestMultiCase(t *testing.T) {
	got := Processor("this is so exciting (up, 2)")
	want := "this is SO EXCITING"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPunctuation(t *testing.T) {
	got := Processor("Hello , world !")
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestQuotes(t *testing.T) {
	got := Processor("I am ' awesome '")
	want := "I am 'awesome'"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAtoAn(t *testing.T) {
	got := Processor("a apple")
	want := "an apple"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
