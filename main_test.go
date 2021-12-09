package main

import "testing"

// 猫😁会瞎几把叫
func TestCat(t *testing.T) {
	saying := Cat()
	if saying != "Miao~~~~" {
		t.Errorf("Cat say %s", saying)
	}
}
