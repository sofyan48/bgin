package moduls

import (
	"github.com/jinzhu/gorm"
)

type LoginScheme struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

func MigrateScheme(db *gorm.DB) {
	db.AutoMigrate(&LoginScheme{})
}
