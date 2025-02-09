package models

type TgAuthData struct {
	Id        int    `json:"id"         db:"tg_user_id" validate:"required"`
	AuthDate  int    `json:"auth_date"  db:"auth_date"  validate:"required"`
	FirstName string `json:"first_name" db:"first_name" validate:"required"`
	LastName  string `json:"last_name"  db:"last_name"`
	Username  string `json:"username"   db:"username"   validate:"required"`
	PhotoURL  string `json:"photo_url"  db:"photo_url"  validate:"required"`
	Hash      string `json:"hash"       db:"hash"       validate:"required"`
}

// TODO: Implement VK auth
type VkAuthData struct {
	VkId        string `json:"id" validate:"required"`
	VkFirstName string `json:"first_name" validate:"required"`
	VkLastName  string `json:"last_name" validate:"required"`
	VkUsername  string `json:"username" validate:"required"`
	VkPhotoURL  string `json:"photo_url" validate:"required"`
	VkHash      string `json:"hash" validate:"required"`
}
