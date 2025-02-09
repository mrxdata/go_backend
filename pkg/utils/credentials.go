package utils

import (
	"fmt"

	"flutty_messenger/pkg/repository"
)

func GetCredentialsByRole(role int) ([]string, error) {
	var credentials []string

	switch role {
	case repository.RoleAdmin: // TO DO: admin creds
		credentials = []string{
			repository.ChatCreateCredential,
			repository.ChatUpdateCredential,
			repository.ChatDeleteCredential,
		}
	case repository.RoleUser:
		credentials = []string{
			repository.ChatCreateCredential, // TO DO: user creds
		}
	default:
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
