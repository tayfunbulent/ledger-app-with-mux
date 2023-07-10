package wallets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"errors"
	"time"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

func GetWalletBalanceAtTime(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestData struct {
			UserID string `json:"userID"`
			Time   string `json:"time"`
		}
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}

		time, err := time.Parse(time.RFC3339, requestData.Time)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid time format")
			return
		}

		balance, err := getBalanceAtTime(db, requestData.UserID, time)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to retrieve balance: %v", err)
			return
		}

		responseData := struct {
			Balance float64 `json:"balance"`
		}{
			Balance: balance,
		}

		jsonData, err := json.Marshal(responseData)
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

func getBalanceAtTime(db *sql.DB, userID string, time time.Time) (float64, error) {
	var balance float64

	err := db.QueryRow(`
		SELECT balance FROM transactions
		WHERE user_id = $1 AND created_at <= $2
		ORDER BY created_at DESC
		LIMIT 1`,
		userID, time).Scan(&balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}
