package models

import "gorm.io/gorm"

type AllowedUser struct {
	gorm.Model
	AuthorizedUser string
	KeyHash        string
	Salt           string
}
