package moduls

import (
	"github.com/jinzhu/gorm"
)

// LoginScheme Field Table loginscheme
type LoginScheme struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

// MigrateScheme Migrate All Scheme
func MigrateScheme(db *gorm.DB) {
	db.AutoMigrate(&LoginScheme{})
}
