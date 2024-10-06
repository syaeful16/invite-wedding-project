package models

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ModelDefault struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

var DB *gorm.DB

func ConnectDB() error {
	mysqlConf := "root:@tcp(127.0.0.1:3306)/invite_wedd?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(mysqlConf), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}

	DB = db

	return nil
}
