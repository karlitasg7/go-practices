package greetings

import (
	"fmt"
	"testing"
)

func TestHelloName(t *testing.T) {

	name := "Karla"
	want := fmt.Sprintf("Hi, %v. Welcome!", name)

	got, err := Hello(name)
	if got != want || err != nil {
		t.Fatalf("Hello(%v) = %q, %v, want %q, nil", name, got, err, want)
	}

}

func TestHelloEmpty(t *testing.T) {

	name := ""
	want := ""

	got, err := Hello(name)
	if got != want || err == nil {
		t.Fatalf("Hello(%v) = %q, %v, want %q, nil", name, got, err, want)
	}
}
