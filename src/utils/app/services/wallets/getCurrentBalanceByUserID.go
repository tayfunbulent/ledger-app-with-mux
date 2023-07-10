package wallets

import (
	"database/sql"
)

func GetCurrentBalanceByUserID(db *sql.DB, userID int) (float64, error) {
	var balance float64
	err := db.QueryRow("SELECT balance FROM wallets WHERE user_id = $1", userID).Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
