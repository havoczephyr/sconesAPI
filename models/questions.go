package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Body     string
	Author   string
	Status   string `gorm:"column:status"`
	Priority bool
}

/* Status can be:
Pending
Denied
Active
Used
*/
