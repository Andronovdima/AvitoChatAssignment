package models

import (
	"time"
)

type Chat struct {
	ID        int64     `json:"id" `
	Name      string    `json:"name" `
	Users     []User    `json:"users" `
	CreatedAt time.Time `json:"created_at" `
}
