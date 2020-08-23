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

func (c *ChatRepository) IsUserInChat(chatID int64, userID string) bool {
	row := c.db.QueryRow(
		"SELECT chat_id "+
			"FROM chat_users "+
			"WHERE chat_id = $1 AND user_id = $2",
		chatID,
		userID,
	)
	if row.Scan(&chatID) != nil {
		return false
	}

	return true
}

func (c *ChatRepository) GetList(userID string) ([]models.Chat, error) {
	var chats []models.Chat
	rows, err := c.db.Query(
			"SELECT chat.id, name, chat.created_at "+
			"FROM chat_users AS c "+
			"LEFT JOIN chat "+
			"ON c.chat_id = chat.id "+
			"LEFT JOIN messages "+
			"ON c.chat_id = messages.chat_id "+
			"WHERE c.user_id = $1 "+
			"ORDER BY messages.created_At DESC",
		userID,
	)

	if err != nil {
		return nil, err
	}

	var ids []int64
	var isAdd bool

	for rows.Next() {
		ch := models.Chat{}
		err := rows.Scan(&ch.ID, &ch.Name, &ch.CreatedAt)
		if err != nil {
			return nil, err
		}

		for _, i := range ids {
			if i == ch.ID {
				isAdd = true
			}
		}

		if !isAdd {
			rows2, err := c.db.Query(
				"SELECT user_id "+
					"FROM chat_users "+
					"WHERE chat_id = $1 ",
				ch.ID,
			)

			if err != nil {
				return nil, err
			}

			for rows2.Next() {
				var s string
				err := rows2.Scan(&s)
				if err != nil {
					return nil, err
				}
				ch.UsersID = append(ch.UsersID, s)
			}

			chats = append(chats, ch)
			ids = append(ids, ch.ID)
		}
		isAdd = false
	}

	return chats, nil
}
