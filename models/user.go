package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255); not null" json:"username" validate:"required"`
	Email    string `gorm:"type:varchar(255); not null; unique" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(255); not null" json:"password" validate:"required,min=8"`
	ModelDefault
}
