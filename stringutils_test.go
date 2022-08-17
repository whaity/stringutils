package stringutils

import "testing"

func TestUpper(t *testing.T) {
	s := Upper("name")
	if s != "NAME" {
		t.Errorf("Wanted NAME got %s", s)
	}
}

func TestLower(t *testing.T) {
	s := Lower("NAME")
	if s != "name" {
		t.Errorf("Wanted name got %s", s)
	}
}
