package models

import "gorm.io/gorm"

type Accounts struct {
	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

func MigrateAccounts(db *gorm.DB) error {
	err := db.AutoMigrate(&Accounts{})
	return err
}