package store

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"svc-user/internal/entity"
)

func NewSQLLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect database")
		return nil
	}

	//auto migrate user entity
	db.AutoMigrate(&entity.User{})
	return db
}
