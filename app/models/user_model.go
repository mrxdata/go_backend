package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	SignedViaTg int = 0
	SignedViaVk int = 1
)

type User struct {
	ID          uuid.UUID `json:"id"            db:"id"            validate:"required,uuid"`
	CreatedAt   time.Time `json:"created_at"    db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"    db:"updated_at"`
	Status      int       `json:"status"        db:"status"        validate:"required,number"`
	Role        int       `json:"role"          db:"role"          validate:"required,number"`
	SignedVia   int       `json:"signed_via"    db:"signed_via"    validate:"required,number"`
	TgUserId    int       `json:"tg_user_id"    db:"tg_user_id"    validate:"omitempty"`
	TgAuthDate  int       `json:"tg_auth_date"  db:"tg_auth_date"  validate:"omitempty"`
	TgFirstName string    `json:"tg_first_name" db:"tg_first_name" validate:"omitempty"`
	TgLastName  string    `json:"tg_last_name"  db:"tg_last_name"`
	TgUsername  string    `json:"tg_username"   db:"tg_username"   validate:"omitempty"`
	TgPhotoURL  string    `json:"tg_photo_url"  db:"tg_photo_url"  validate:"omitempty"`
	TgHash      string    `json:"tg_hash"       db:"tg_hash"       validate:"omitempty"`
}

func (user *User) UpdateUserTgInfo(tg TgAuthData) {
	user.TgUserId = tg.Id
	user.TgAuthDate = tg.AuthDate
	user.TgFirstName = tg.FirstName
	user.TgLastName = tg.LastName
	user.TgUsername = tg.Username
	user.TgPhotoURL = tg.PhotoURL
	user.TgHash = tg.Hash
}
