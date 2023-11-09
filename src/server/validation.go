package server

import (
	"errors"
	"strconv"
)

var (
	vALID_CODES = [8]uint16{2320, 1121, 2322, 2435, 1036, 1949, 1606, 1559}
)

const (
	iD_MASK = 7
)

func extractCode(message string) (uint16, error) {
	if len(message) != 5 {
		return 0, errors.New("empty message")
	}
	digits := message[0:2] + message[3:5]
	code, err := strconv.Atoi(digits)
	if err != nil {
		return 0, err
	}
	return uint16(code), nil
}

func validateCode(code string) (uint8, error) {
	id, err := extractCode(code)
	if err != nil {
		return 0, err
	}
	//check if id is in valid group
	{
		i := 0
		isValid := false
		for i < len(vALID_CODES) && !isValid {
			isValid = uint16(id) == vALID_CODES[i]
			i++
		}

		if !isValid {
			err = errors.New("invalid id")
		}
	}
	return uint8(id & iD_MASK), err
}
