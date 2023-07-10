package users

import (
	"database/sql"
)

func GetUserIDByCredentials(db *sql.DB, email, password string) (int, error) {
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE email = $1 AND password = $2", email, password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
