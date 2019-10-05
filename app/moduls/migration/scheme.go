package moduls

import (
	"github.com/jinzhu/gorm"
)

type LoginScheme struct {
	ID       int     `json:"id,omitempty"`
	Username *string `json:"username" gorm:"not null"`
	Password *string `json:"password" gorm:"not null"`
}

type Userdata struct {
	ID      int     `json:"id,omitempty"`
	Name    *string `json:"name" gorm:"not null"`
	Address *string `json:"address" gorm:"not null"`
}

func MigrateScheme(db *gorm.DB) {
	db.AutoMigrate(&LoginScheme{}, &Userdata{})
}
