package database

import (
	"database/sql"
	"fmt"
)

func Migrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id varchar(255) NOT NULL,
		price float NOT NULL,
		tax float NOT NULL,
		final_price float NOT NULL,
		PRIMARY KEY (id)
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating orders table: %w", err)
	}

	fmt.Println("Migration completed: Table 'orders' is ready.")
	return nil
}