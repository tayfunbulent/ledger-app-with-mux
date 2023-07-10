package wallets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"ledgerApp/src/utils/app/services/users"
)

func GetWallet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.Header.Get("email")
		password := r.Header.Get("password")

		userID, err := users.GetUserIDByCredentials(db, email, password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid credentials")
			return
		}

		wallet, err := GetWalletByUserID(db, userID)
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
