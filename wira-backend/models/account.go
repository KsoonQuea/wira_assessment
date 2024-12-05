package models

import (
	"wira-backend/connection"

	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Account struct {
	ID        	uint  		`gorm:"primary key;autoIncrement" json:"acc_id"`
	Username    *string 	`json:"username"`
	Email     	*string 	`json:"email"`
	Password	*string		`json:"encrypted_password"`
	SecretKey 	*string 	`json:"secretkey_2fa"`
}

func stringPtr(s string) *string {
    return &s
}


func GetAccounts() ([]Account, error) {
	results, err := connection.GetAll("accounts")

	if err != nil {
		log.Printf("Error querying accounts: %v", err)
		return nil, err
	}

	var collections []Account
	for _, result := range results {

		fmt.Println(result["username"])

		item := Account{
			ID:       uint(result["acc_id"].(int64)) ,
			Username: stringPtr(result["username"].(string)),
			Email:    stringPtr(result["email"].(string)),
		}
		collections = append(collections, item)
	}
	
	return collections, nil
}