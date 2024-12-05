package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")

	// Create items table if it doesn't exist
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return DB
}

// GetAll retrieves all rows from a specified table
func GetAll(table string) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s", table)
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}

		err := rows.Scan(pointers...)
		if err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, column := range columns {
			row[column] = values[i]
		}
		results = append(results, row)
	}

	return results, nil
}

func GetJoinedData(query string, args ...interface{}) ([]map[string]interface{}, error) {
    rows, err := DB.Query(query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    columns, err := rows.Columns()
    if err != nil {
        return nil, err
    }

    var results []map[string]interface{}

    for rows.Next() {
        values := make([]interface{}, len(columns))
        pointers := make([]interface{}, len(columns))
        for i := range values {
            pointers[i] = &values[i]
        }

        err := rows.Scan(pointers...)
        if err != nil {
            return nil, err
        }

        row := make(map[string]interface{})
        for i, column := range columns {
            row[column] = values[i]
        }
        results = append(results, row)
    }

    return results, nil
}

// // CreateData inserts a new row into the specified table
// func CreateData(table string, data map[string]interface{}) (map[string]interface{}, error) {
// 	columns := make([]string, 0, len(data))
// 	values := make([]interface{}, 0, len(data))
// 	placeholders := make([]string, 0, len(data))

// 	i := 1
// 	for column, value := range data {
// 		columns = append(columns, column)
// 		values = append(values, value)
// 		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
// 		i++
// 	}

// 	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING *",
// 		table,
// 		data.Join(columns, ", "),
// 		data.Join(placeholders, ", "))

// 	row := DB.QueryRow(query, values...)

// 	columns, err := row.Columns()
// 	if err != nil {
// 		return nil, err
// 	}

// 	pointers := make([]interface{}, len(columns))
// 	container := make([]interface{}, len(columns))
// 	for i := range pointers {
// 		pointers[i] = &container[i]
// 	}

// 	err = row.Scan(pointers...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result := make(map[string]interface{})
// 	for i, column := range columns {
// 		result[column] = container[i]
// 	}

// 	return result, nil
// }

