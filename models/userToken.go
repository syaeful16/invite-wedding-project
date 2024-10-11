package models

type UserToken struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token" validate:"required"`
	ModelDefault
}
