package models

type Invitation struct {
	ID             uint    `gorm:"primaryKey" json:"id"`
	InvitationCode string  `gorm:"type:varchar(255); not null" json:"invitation_code" validate:"required"`
	MaleNickname   string  `gorm:"type:varchar(255); not null" json:"male_nickname" validate:"required"`
	FemaleNickname string  `gorm:"type:varchar(255); not null" json:"female_nickname" validate:"required"`
	Status         string  `gorm:"type:varchar(255); not null" json:"status" validate:"required"`
	UserID         uint    `json:"user_id" validate:"required"`
	Bride          []Bride `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	ModelDefault
}
