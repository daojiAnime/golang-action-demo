package main

import "testing"

// η«πδΌηε ζε«
func TestCat(t *testing.T) {
	saying := Cat()
	if saying != "Miao~~~~" {
		t.Errorf("Cat say %s", saying)
	}
}
