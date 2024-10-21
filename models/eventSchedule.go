package models

type EventSchedule struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	EventName    string `gorm:"varchar(255); not null" validate:"required" json:"event_name"`
	LocationName string `gorm:"varchar(255); not null" validate:"required" json:"location_name"`
	Address      string `gorm:"varchar(255); not null" validate:"required" json:"address"`
	EventDate    string `gorm:"date; not null" validate:"required" json:"event_date"`
	StartTime    string `gorm:"time; not null" validate:"required" json:"start_time"`
	EndTime      string `gorm:"time; not null" json:"end_time"`
	LinkLocation string `gorm:"varchar(255); not null" validate:"required" json:"LinkLocation"`
	InvitationID uint   `json:"invitation_id" validate:"required"`
	ModelDefault
}
