package server

import "testing"

func TestExtractCode(t *testing.T) {
	_, err := extractCode("23:200")
	if err == nil {
		t.Error("too long message not detected")
	}
	_, err = extractCode("23:2")
	if err == nil {
		t.Error("too short message not detected")
	}
	_, err = extractCode("23:2a")
	if err == nil {
		t.Error("non-digit character not detected")
	}
	id, err := extractCode("23:20")
	if err != nil {
		t.Error("valid code not detected")
	}
	if id != 2320 {
		t.Error("invalid code returned")
	}
}

func TestValidateCode(t *testing.T) {
	id, _ := validateCode("23:20")
	if id != 0 {
		t.Error("invalid id returned")
	}
	id, _ = validateCode("11:21")
	if id != 1 {
		t.Error("invalid id returned")
	}
	id, _ = validateCode("23:22")
	if id != 2 {
		t.Error("invalid id returned")
	}
	id, _ = validateCode("23:23")
	if id != 3 {
		t.Error("invalid id returned")
	}
	id, _ = validateCode("10:36")
	if id != 4 {
		t.Error("invalid id returned")
	}
	id, _ = validateCode("19:49")
	if id != 5 {
		t.Error("invalid id returned")
	}
	id, _ = validateCode("16:06")
	if id != 6 {
		t.Error("invalid id returned")
	}
	id, _ = validateCode("15:59")
	if id != 7 {
		t.Error("invalid id returned")
	}
	_, err := validateCode("99:99")
	if err == nil {
		t.Error("invalid code not detected")
	}
}
