package wallets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"ledgerApp/src/utils/models"
)

func GetAllWallets(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, balance, created_at FROM wallets")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to retrieve wallets: %v", err)
			return
		}
		defer rows.Close()

		wallets := []*models.Wallet{}
		for rows.Next() {
			wallet := &models.Wallet{}
			err := rows.Scan(&wallet.ID, &wallet.Balance, &wallet.CreatedAt)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Failed to retrieve wallets: %v", err)
				return
			}
			wallets = append(wallets, wallet)
		}

		jsonData, err := json.Marshal(wallets)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to marshal JSON: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}
