package wallets

import (
	"database/sql"
	"fmt"
)

func GetWalletBalance(db *sql.DB, userID int) (float64, error) {
	var balance float64
	err := db.QueryRow("SELECT balance FROM wallets WHERE user_id = $1", userID).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("wallet not found")
		}
		return 0, err
	}
	return balance, nil
}
