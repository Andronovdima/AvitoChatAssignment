package repository

import (
	"database/sql"
	"github.com/Andronovdima/AvitoChatAssignment/internal/models"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(thisDB *sql.DB) *ChatRepository {
	chatRep := &ChatRepository{
		db: thisDB,
	}
	return chatRep
}

func (c *ChatRepository) Create(chat *models.Chat) error {
	err := c.db.QueryRow(
		"INSERT INTO chat (name) "+
			"VALUES ($1) RETURNING id",
		chat.Name,
	).Scan(&chat.ID)

	if err != nil {
		return err
	}

	for _, id := range chat.UsersID {
		_, err = c.db.Exec(
			"INSERT INTO chat_users (user_id, chat_id) "+
				"VALUES ($1, $2)",
			id,
			chat.ID,
		)
	}

	return err
}

func (c *ChatRepository) IsExist(name string) bool {
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

func (c *ChatRepository) IsExistByID(ID int64) bool {
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
