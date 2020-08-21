package models

import (
	"time"
)

type Chat struct {
	ID        int64     `json:"id" `
	Name      string    `json:"name" `
	UsersID   []string  `json:"users" `
	CreatedAt time.Time `json:"created_at" `
}

type GetChat struct {
	User string `json:"user" `
}
