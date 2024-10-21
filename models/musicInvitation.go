package models

type MusicInvitation struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	SongTitle    string `gorm:"varchar(255); not null" validate:"required" json:"song_title"`
	SongFile     string `gorm:"varchar(255); not null" validate:"required" json:"song_file"`
	ArtistName   string `gorm:"varchar(255); not null" validate:"required" json:"artist_name"`
	InvitationID uint   `json:"invitation_id" validate:"required"`
	ModelDefault
}
