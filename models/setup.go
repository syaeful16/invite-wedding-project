package models

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ModelDefault struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

var DB *gorm.DB

func ConnectDB() {
	mysqlConf := "root:@tcp(localhost:3306)/invite_wedd?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Log to stdout
		logger.Config{
			SlowThreshold:             time.Second, // Log slow SQL queries (default: 200ms)
			LogLevel:                  logger.Info, // Log level (e.g., Info, Warn, Error)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable colorful output
		},
	)

	db, err := gorm.Open(mysql.Open(mysqlConf), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("Error connecting to database")
	}

	if err := db.AutoMigrate(&User{}, &UserToken{}, &Invitation{}, &Bride{}); err != nil {
		panic("Error migrating database")
	}

	DB = db
}
