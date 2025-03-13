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
	if err := user.validate(); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("field Name is required")
	}

	if user.Email == "" {
		return errors.New("field Email is required")
	}

	if user.NickName == "" {
		return errors.New("field NickName is required")
	}

	if user.Password == "" {
		return errors.New("field Password is required")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Password = strings.TrimSpace(user.Password)
}
