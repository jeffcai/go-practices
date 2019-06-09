package main

import (
	"testing"
)

func TestTotalFruit(t *testing.T) {
	tree := []int{0, 1, 2, 2}
	fruits := TotalFruit(tree)
	if fruits != 3 {
		t.Errorf("fruits %d", fruits)
	}

	tree = []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fruits = TotalFruit(tree)
	if fruits != 5 {
		t.Errorf("fruits %d", fruits)
	}

	tree = []int{0}
	fruits = TotalFruit(tree)
	if fruits != 1 {
		t.Errorf("fruits %d", fruits)
	}
}
