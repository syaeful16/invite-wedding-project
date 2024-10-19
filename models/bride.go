package models

type Bride struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"varchar(255); not null" validate:"required" json:"name"`
	Photo         string `gorm:"varchar(255); not null" validate:"required" json:"photo"`
	FatherName    string `gorm:"varchar(255); not null" validate:"required" json:"father_name"`
	MotherName    string `gorm:"varchar(255); not null" validate:"required" json:"mother_name"`
	InstagramLink string `gorm:"varchar(255)" json:"instagram_link"`
	FacebookLink  string `gorm:"varchar(255)" json:"facebook_link"`
	InvitationID  uint   `json:"invitation_id" validate:"required"`
	ModelDefault
}
