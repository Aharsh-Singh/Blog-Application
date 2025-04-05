package models

import "gorm.io/gorm"

var DB *gorm.DB

type User struct {
	gorm.Model
	Name     string
	Supabase_ID string
	Email    string `gorm:"unique"`
}