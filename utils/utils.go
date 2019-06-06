package utils

import (
	 "github.com/nu7hatch/gouuid"
)

func GenerateUUID() (string, error) {
	id, err := uuid.NewV4()
	return id.String(), err
}