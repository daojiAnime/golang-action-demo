package main

import "testing"

// 猫会瞎几把叫
func TestCat(t *testing.T) {
	saying := Cat()
	if saying != "Miao~~~~" {
		t.Errorf("Cat say %s", saying)
	}
}
