package models

import "time"

type User struct {
	ID        int64     `json:"id.omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	NickName  string    `json:"nick"`
	Password  string    `json:"password.omitempty"`
	CreatedAt time.Time `json:"created_at.omitempty"`
}
