package server

import (
	"errors"
)

var (
	vALID_CODES = [PEOPLE_COUNT]string{"Elf", "Gingerbread", "Santa", "Mrs. Clause", "Raindeer", "Fairy lights", "Carol", "Tinsel", "Mistletoe"}
)

func validateCode(code string) (uint8, error) {
	i := 0
	isValid := false
	for i < len(vALID_CODES) && !isValid {
		isValid = code == vALID_CODES[i]
		i++
	}

	if !isValid {
		return 0, errors.New("invalid id")
	}
	return uint8(i - 1), nil
}
