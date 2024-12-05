package models

import (
	"wira-backend/connection"

	_ "github.com/lib/pq"
)

type Character struct {
	ID        	uint  		`gorm:"primary key;autoIncrement" json:"char_id"`
	Account    	*string 	`json:"acc_id"`
	Class     	*string 	`json:"class_id"`
}

func GetCharacters() ([]Character, error) {
	rows, err := connection.DB.Query("SELECT * FROM characters;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var model_arrs []Character
	for rows.Next() {
		var arr_item Character
		if err := rows.Scan(&arr_item.ID, &arr_item.Account, &arr_item.Class); err != nil {
			return nil, err
		}
		model_arrs = append(model_arrs, arr_item)
	}
	return model_arrs, nil
}

// func AddItem(name string) (Item, error) {
// 	var item Item
// 	err := db.GetDB().QueryRow("INSERT INTO items (name) VALUES ($1) RETURNING id, name", name).Scan(&item.ID, &item.Name)
// 	if err != nil {
// 		return Item{}, err
// 	}
// 	return item, nil
// }

// func UpdateItem(id int, name string) (Item, error) {
// 	var item Item
// 	err := db.GetDB().QueryRow("UPDATE items SET name = $1 WHERE id = $2 RETURNING id, name", name, id).Scan(&item.ID, &item.Name)
// 	if err != nil {
// 		return Item{}, err
// 	}
// 	return item, nil
// }

// func DeleteItem(id int) error {
// 	result, err := db.GetDB().Exec("DELETE FROM items WHERE id = $1", id)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if rowsAffected == 0 {
// 		return fmt.Errorf("item with id %d not found", id)
// 	}
// 	return nil
// }

// type Account struct {
// 	AccID    uint   `gorm:"column:acc_id;primaryKey"`
// 	Username string `gorm:"column:username"`
// 	Email    string `gorm:"column:email"`
// }

// func MigrateAccounts(db *gorm.DB) error {
// 	err := db.AutoMigrate(&Accounts{})
// 	return err
// }