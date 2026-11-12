package genz_test

import (
	"testing"

	"github.com/brunty/genz"
)

func TestItsGiving(t *testing.T) {
	if got, want := genz.ItsGiving("main character"), "it's giving main character"; got != want {
		t.Errorf("ItsGiving(%q) = %q, want %q", "main character", got, want)
	}
}

func TestEra(t *testing.T) {
	if got, want := genz.Era("gym"), "in my gym era"; got != want {
		t.Errorf("Era(%q) = %q, want %q", "gym", got, want)
	}
}

func TestRentFree(t *testing.T) {
	if got, want := genz.RentFree("that song"), "that song is living rent free in my head"; got != want {
		t.Errorf("RentFree(%q) = %q, want %q", "that song", got, want)
	}
}
