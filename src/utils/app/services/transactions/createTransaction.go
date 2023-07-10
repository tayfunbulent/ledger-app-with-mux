package transactions

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"errors"
	"ledgerApp/src/utils/app/services/users"
	"ledgerApp/src/utils/app/services/wallets"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

func CreateTransaction(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		senderEmail := r.Header.Get("email")
		senderPassword := r.Header.Get("password")

		senderID, err := users.GetUserIDByCredentials(db, senderEmail, senderPassword)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid credentials")
			return
		}

		var transferData struct {
			RecipientEmail string  `json:"recipient_email"`
			Amount         float64 `json:"amount"`
		}
		err = json.NewDecoder(r.Body).Decode(&transferData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}

		recipientID, err := users.GetUserIDByEmail(db, transferData.RecipientEmail)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Recipient not found")
			return
		}

		err = createTransaction(db, senderID, recipientID, transferData.Amount)
		if err != nil {
			if err == ErrInsufficientFunds {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Insufficient funds")
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Failed to transfer funds: %v", err)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Funds transferred successfully")
	}
}

func createTransaction(db *sql.DB, senderID, recipientID int, amount float64) error {
	senderBalance, err := wallets.GetWalletBalance(db, senderID)
	if err != nil {
		return err
	}

	if amount > senderBalance {
		return ErrInsufficientFunds
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE wallets SET balance = balance - $1 WHERE user_id = $2", amount, senderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE wallets SET balance = balance + $1 WHERE user_id = $2", amount, recipientID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO transactions (user_id, description, amount) VALUES ($1, $2, $3)",
		senderID, "Transfer to "+strconv.Itoa(recipientID), -amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO transactions (user_id, description, amount) VALUES ($1, $2, $3)",
		recipientID, "Transfer from "+strconv.Itoa(senderID), amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
