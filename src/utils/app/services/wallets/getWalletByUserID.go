package wallets

import (
	"database/sql"
	"ledgerApp/src/utils/models"
)

func GetWalletByUserID(db *sql.DB, userID int) (*models.Wallet, error) {
	wallet := &models.Wallet{}
	err := db.QueryRow("SELECT id, user_id, balance, created_at FROM wallets WHERE user_id = $1", userID).Scan(&wallet.ID, &wallet.UserID, &wallet.Balance, &wallet.CreatedAt)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}
