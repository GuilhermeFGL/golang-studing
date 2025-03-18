package models

import (
	"errors"
	"example.com/m/v2/API/src/security"
	"github.com/badoux/checkmail"
	"log"
	"strings"
	"time"
)

type User struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	NickName  string    `json:"nickname"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare validate and format user data
func (user *User) Prepare() error {
	if err := user.validate(true); err != nil {
		return err
	}

	return user.Format()
}

func (user *User) PrepareForUpdate() error {
	if err := user.validate(false); err != nil {
		return err
	}

	return user.Format()
}

func (user *User) validate(creating bool) error {
	if user.Name == "" {
		return errors.New("field Name is required")
	}

	if user.Email == "" {
		return errors.New("field Email is required")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("field Email is invalid: " + err.Error())
	}

	if user.NickName == "" {
		return errors.New("field NickName is required")
	}

	if creating && user.Password == "" {
		return errors.New("field Password is required")
	}

	return nil
}

func (user *User) Format() error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.NickName = strings.TrimSpace(user.NickName)

	if user.Password != "" {
		user.Password = strings.TrimSpace(user.Password)

		hash, err := security.Hash(user.Password)
		if err != nil {
			log.Fatal(err)
			return err
		}
		user.Password = string(hash)
	}

	return nil
}
