package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"ledgerApp/src/utils/models"
)

func GetAllUsersWithWallet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT users.id, users.username, users.email, users.password,
			wallets.id, wallets.user_id, wallets.balance, wallets.created_at
			FROM users
			INNER JOIN wallets ON users.id = wallets.user_id
		`)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to retrieve users: %v", err)
			return
		}
		defer rows.Close()

		users := []*models.UserWithWallet{}
		for rows.Next() {
			user := &models.UserWithWallet{
				User:   &models.User{},
				Wallet: &models.Wallet{},
			}
			err := rows.Scan(
				&user.User.ID, &user.User.Username, &user.User.Email, &user.User.Password,
				&user.Wallet.ID, &user.Wallet.UserID, &user.Wallet.Balance, &user.Wallet.CreatedAt,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Failed to retrieve users: %v", err)
				return
			}
			users = append(users, user)
		}

		jsonData, err := json.Marshal(users)
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
