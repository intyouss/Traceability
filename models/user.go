package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(64);not null;unique" json:"username"`
	Password string `gorm:"type:varchar(128);not null" json:"password"`
	Avatar   string `gorm:"type:varchar(256)" json:"avatar"`
	Email    string `gorm:"type:varchar(128);unique" json:"email"`
	Mobile   string `gorm:"type:varchar(11);unique" json:"mobile"`
}
