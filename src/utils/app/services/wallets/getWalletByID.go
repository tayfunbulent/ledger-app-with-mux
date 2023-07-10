package wallets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"ledgerApp/src/utils/models"
)

func GetWalletByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		walletID, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid wallet ID")
			return
		}

		wallet, err := getWalletByID(db, walletID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to retrieve wallet: %v", err)
			return
		}

		jsonData, err := json.Marshal(wallet)
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

func getWalletByID(db *sql.DB, walletID int) (*models.Wallet, error) {
	row := db.QueryRow("SELECT id, balance, created_at FROM wallets WHERE id = $1", walletID)

	wallet := &models.Wallet{}
	err := row.Scan(&wallet.ID, &wallet.Balance, &wallet.CreatedAt)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
