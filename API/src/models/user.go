package models

import (
	"errors"
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

	user.Format()
	return nil
}

func (user *User) PrepareForUpdate() error {
	if err := user.validate(false); err != nil {
		return err
	}

	user.Format()
	return nil
}

func (user *User) validate(creating bool) error {
	if user.Name == "" {
		return errors.New("field Name is required")
	}

	if user.Email == "" {
		return errors.New("field Email is required")
	}

	if user.NickName == "" {
		return errors.New("field NickName is required")
	}

	if creating && user.Password == "" {
		return errors.New("field Password is required")
	}

	return nil
}

func (user *User) Format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Password = strings.TrimSpace(user.Password)
}
