package wallets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"math"
	"errors"
	"ledgerApp/src/utils/app/services/users"
)

func UpdateWalletBalance(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.Header.Get("email")
		password := r.Header.Get("password")

		userID, err := users.GetUserIDByCredentials(db, email, password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid credentials")
			return
		}

		var updateData struct {
			Amount float64 `json:"amount"`
		}

		err = json.NewDecoder(r.Body).Decode(&updateData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}

		err = updateWalletBalance(db, userID, updateData.Amount)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to update wallet balance: %v", err)
			return
		}

		err = createTransaction(db, userID, updateData.Amount)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to create transaction: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Wallet balance updated successfully")
	}
}

func updateWalletBalance(db *sql.DB, userID int, amount float64) error {
	currentBalance, err := GetCurrentBalanceByUserID(db, userID)
	if err != nil {
		return err
	}

	if amount < 0 && math.Abs(amount) > currentBalance {
		return errors.New("insufficient funds")
	}

	_, err = db.Exec("UPDATE wallets SET balance = balance + $1 WHERE user_id = $2", amount, userID)
	if err != nil {
		return err
	}

	return nil
}

func createTransaction(db *sql.DB, userID int, amount float64) error {
	_, err := db.Exec("INSERT INTO transactions (user_id, description, amount) VALUES ($1, $2, $3)",
		userID, "Wallet update", amount)
	if err != nil {
		return err
	}

	return nil
}
