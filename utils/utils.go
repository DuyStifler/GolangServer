package utils

import (
	 "github.com/nu7hatch/gouuid"
)

func GenerateUUID() (string, error) {
	id, err := uuid.NewV4()
	return id.String(), err
}

func InStringArray(val string, arr []string) bool {
	for _, each := range arr {
		if val == each {
			return true
		}
	}

	return false
}