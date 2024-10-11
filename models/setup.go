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

func ConnectDB() {
	mysqlConf := "root:@tcp(localhost:3306)/invite_wedd?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(mysqlConf), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	}

	if err := db.AutoMigrate(&User{}, &UserToken{}); err != nil {
		panic("Error migrating database")
	}

	DB = db
}
