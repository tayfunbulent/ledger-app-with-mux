package transactions

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"ledgerApp/src/utils/models"
	"ledgerApp/src/utils/app/services/users"
)

func GetTransactions(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.Header.Get("email")
		password := r.Header.Get("password")

		userID, err := users.GetUserIDByCredentials(db, email, password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid credentials")
			return
		}

		transactions, err := getTransactions(db, userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to retrieve user transactions: %v", err)
			return
		}

		jsonData, err := json.Marshal(transactions)
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

func getTransactions(db *sql.DB, userID int) ([]models.Transaction, error) {
	rows, err := db.Query("SELECT id, user_id, description, amount, balance, created_at FROM transactions WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []models.Transaction{}

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Description, &transaction.Balance, &transaction.Amount, &transaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
