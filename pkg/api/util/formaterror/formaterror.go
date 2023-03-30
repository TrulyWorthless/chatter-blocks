package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "uuid") {
		return errors.New("uuid already taken")
	}

	if strings.Contains(err, "alias") {
		return errors.New("alias already taken")
	}

	if strings.Contains(err, "address") {
		return errors.New("address already taken")
	}

	if strings.Contains(err, "publickey") {
		return errors.New("publickey already taken")
	}

	if strings.Contains(err, "channel") {
		return errors.New("channel already taken")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("incorrect password")
	}

	return errors.New("incorrect details")
}
