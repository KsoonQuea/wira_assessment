package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
	"fmt"
	// "github.com/goravel/framework/support/facades"
)

type Account struct {
	orm.Model
	AccID    uint   `gorm:"column:acc_id;primaryKey"`
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
}

// TableName specifies the database table name
func (Account) TableName() string {
	return "account"
}

// Repository for Account-related database operations
type AccountRepository struct {}

// Index method to list all accounts
func (r *AccountRepository) Index() ([]Account, error) {
	var accounts []Account

	// Retrieve all accounts from the database
	result := facades.Orm().Query().Table("account")
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	if result == nil {
		return nil, fmt.Errorf("query failed")
	}
	
	return accounts, nil
}