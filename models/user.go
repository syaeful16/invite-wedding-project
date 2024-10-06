package models

type User struct {
	Username string `gorm:"type:varchar(255); not null: unique" json:"username"`
	Email    string `gorm:"type:varchar(255); not null: unique" json:"email"`
	Password string `gorm:"type:varchar(255); not null" json:"password"`
	ModelDefault
}
