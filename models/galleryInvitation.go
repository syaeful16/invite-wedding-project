package models

type GalleryInvitation struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Photo        string `gorm:"varchar(255); not null" validate:"required" json:"photo"`
	InvitationID uint   `json:"invitation_id" validate:"required"`
	ModelDefault
}
