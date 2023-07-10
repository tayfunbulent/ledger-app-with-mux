package users

import (
	"database/sql"
	"fmt"
)

func GetUserIDByEmail(db *sql.DB, email string) (int, error) {
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE email = $1", email).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("user not found")
		}
		return 0, err
	}
	return userID, nil
}
