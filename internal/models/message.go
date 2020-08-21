package models

import "time"

type Message struct {
	ID        int64     `json:"id" `
	ChatID    string    `json:"chat" `
	AuthorID  int64     `json:"author" `
	Text      string    `json:"text" `
	CreatedAt time.Time `json:"created_at" `
}
