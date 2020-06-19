package hashing

import "testing"

// test the block hash function
func TestHashIsSame(t *testing.T) {
	input := "Test"
	t1 := Hash(input)
	t2 := Hash(input)

	if t1 != t2 {
		t.Errorf("identical inputs should have matching hash : %v != %v", t1, t2)
	}
}

func TestHashIsDifferent(t *testing.T) {
	t1 := Hash("Test")
	t2 := Hash("Test2")

	if t1 == t2 {
		t.Errorf("Hashes with different inputs should produce different outputs")
	}
}
