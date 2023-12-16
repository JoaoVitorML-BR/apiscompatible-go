package validation

import (
	"errors"
	"login/domain/secure"
	"login/infra/mysql/bridge"
	"strings"
)

type ValidStruct struct {
	UserParams interface{}
}

func (user *ValidStruct) Prepare(etap string) error {
	if err := user.validate(etap); err != nil {
		return err
	}

	if err := user.format(etap); err != nil {
		return errors.New("Failed to format user data")
	}

	return nil
}

func (user *ValidStruct) validate(etap string) error {
    if params, ok := user.UserParams.(*bridge.CreateUserParams); ok {
        if params.Name == "" {
            return errors.New("Name's invalid")
        }
        if etap == "createNewUser" && len(params.Password) == 0 {
            return errors.New("Password's invalid")
        }
    } else if _, ok := user.UserParams.(*bridge.UpdateUserParams); ok {
        // Validation for UpdateUserParams if needed
    } else {
        return errors.New("Invalid user params type")
    }

    return nil
}

func (user *ValidStruct) format(etap string) error {
    switch params := user.UserParams.(type) {
    case *bridge.CreateUserParams:
        if etap == "createNewUser" {
            params.Name = strings.TrimSpace(params.Name)
            // passing password to hash
            hashedPassword, err := secure.Hash(string(params.Password))
            if err != nil {
                return errors.New("Failed to hash password")
            }
            params.Password = string(hashedPassword)
        }
    case *bridge.UpdateUserParams:
        params.Name = strings.TrimSpace(params.Name)
    }

    return nil
}
