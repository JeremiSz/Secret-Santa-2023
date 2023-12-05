package server

import "testing"

func TestValidateCode(t *testing.T) {
	id, _ := validateCode("Elf")
	if id != 0 {
		t.Error("invalid id returned", id, "expected", 0)
	}
	id, _ = validateCode("Gingerbread")
	if id != 1 {
		t.Error("invalid id returned", id, "expected", 1)
	}
	id, _ = validateCode("Santa")
	if id != 2 {
		t.Error("invalid id returned", id, "expected", 2)
	}
	id, _ = validateCode("Mrs. Clause")
	if id != 3 {
		t.Error("invalid id returned", id, "expected", 3)
	}
	id, _ = validateCode("Raindeer")
	if id != 4 {
		t.Error("invalid id returned", id, "expected", 4)
	}
	id, _ = validateCode("Fairy lights")
	if id != 5 {
		t.Error("invalid id returned", id, "expected", 5)
	}
	id, _ = validateCode("Carol")
	if id != 6 {
		t.Error("invalid id returned", id, "expected", 6)
	}
	id, _ = validateCode("Tinsel")
	if id != 7 {
		t.Error("invalid id returned", id, "expected", 7)
	}
	id, _ = validateCode("Mistletoe")
	if id != 8 {
		t.Error("invalid id returned", id, "expected", 8)
	}
}
