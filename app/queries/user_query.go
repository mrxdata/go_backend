package queries

import (
	"flutty_messenger/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) GetUserRoleByID(id uuid.UUID) (int, error) {
	var role int
	query := "SELECT role FROM users WHERE id=$1"

	err := q.Get(&role, query, id)
	if err != nil {
		return -1, err
	}

	return role, nil
}

func (q *UserQueries) GetUserByID(id uuid.UUID) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM users WHERE id = $1`

	err := q.Get(&user, query, id)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	return user, nil
}

func (q *UserQueries) GetUserByTgUserID(tgUserID int) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM users WHERE tg_user_id = $1`

	err := q.Get(&user, query, tgUserID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (q *UserQueries) CreateUser(user *models.User) (uuid.UUID, error) {
	var generatedID uuid.UUID
	var query string

	if user.SignedVia == models.SignedViaTg {
		query = `INSERT INTO users (
    signed_via,
    tg_user_id,
    tg_first_name,
    tg_last_name,
    tg_username,
    tg_photo_url,
    tg_hash
) VALUES (
    $1,           -- signed_via (int)
    $2,           -- tg_user_id (int)
    $3,           -- tg_first_name (string)
    $4,           -- tg_last_name (string)
    $5,           -- tg_username (string)
    $6,           -- tg_photo_url (string)
    $7            -- tg_hash (string)
) RETURNING id;
`
	}

	err := q.Get(&generatedID, query,
		user.SignedVia,
		user.TgUserId,
		user.TgFirstName,
		user.TgLastName,
		user.TgUsername,
		user.TgPhotoURL,
		user.TgHash,
	)

	if err != nil {
		return uuid.Nil, err
	}

	return generatedID, nil
}
