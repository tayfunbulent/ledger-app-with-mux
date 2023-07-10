package wallets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"math"
)

func UpdateWalletBalanceByUserID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateData struct {
			UserID string  `json:"id"`
			Amount float64 `json:"amount"`
		}

		err := json.NewDecoder(r.Body).Decode(&updateData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}

		userID, err := strconv.Atoi(updateData.UserID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Invalid user ID")
			return
		}

		err = updateWalletBalanceByUserID(db, userID, updateData.Amount)
		if err != nil {
			if err == ErrInsufficientFunds {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			fmt.Fprintf(w, "Failed to update wallet balance: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Wallet balance updated successfully")
	}
}

func updateWalletBalanceByUserID(db *sql.DB, userID int, amount float64) error {
	currentBalance, err := GetCurrentBalanceByUserID(db, userID)
	if err != nil {
		return err
	}

	if amount < 0 && math.Abs(amount) > currentBalance {
		return ErrInsufficientFunds
	}

	_, err = db.Exec("UPDATE wallets SET balance = balance + $1 WHERE user_id = $2", amount, userID)
	if err != nil {
		return err
	}

	return nil
}
