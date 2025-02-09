package utils

import (
	"fmt"

	"flutty_messenger/pkg/repository"
)

func VerifyRole(role int) (int, error) {
	switch role {
	case repository.RoleUser:
	case repository.RoleAdmin:
	default:
		return -1, fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
