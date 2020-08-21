package repository

import (
	"database/sql"
	"github.com/Andronovdima/AvitoChatAssignment/internal/models"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(thisDB *sql.DB) *MessageRepository {
	messageRep := &MessageRepository{
		db: thisDB,
	}
	return messageRep
}

func (c *MessageRepository) Create(message *models.Message) error {
	return c.db.QueryRow(
		"INSERT INTO messages (chat_id, user_id, text) "+
			"VALUES ($1, $2, $3) RETURNING id",
		message.ChatID,
		message.AuthorID,
		message.Text,
	).Scan(&message.ID)
}

func (c *MessageRepository) IsExist(name string) bool {
	row := c.db.QueryRow(
		"SELECT name "+
			"FROM chat "+
			"WHERE name = $1",
		name,
	)
	if row.Scan(&name) != nil {
		return false
	}

	return true

}

func (c *MessageRepository) IsExistByID(ID int64) bool {
	row := c.db.QueryRow(
		"SELECT id "+
			"FROM chat "+
			"WHERE id = $1",
		ID,
	)
	if row.Scan(&ID) != nil {
		return false
	}

	return true

}
