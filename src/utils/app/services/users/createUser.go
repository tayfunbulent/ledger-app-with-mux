package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unicode"
	"ledgerApp/src/utils/models"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Failed to decode user data: %v", err)
			return
		}

		if user.Username == "" || user.Email == "" || user.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Missing required fields")
			return
		}

		if !isEmailValid(user.Email) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid email format")
			return
		}

		if !isPasswordValid(user.Password) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid password format")
			return
		}

		err = createUserWithWallet(db, &user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to create user: %v", err)
			return
		}

		jsonData, err := json.Marshal(user)
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

func createUserWithWallet(db *sql.DB, user *models.User) error {
	if isEmailTaken(db, user.Email) {
		return fmt.Errorf("email is already taken")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = tx.QueryRow("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO wallets (user_id) VALUES ($1)", user.ID)
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

func isEmailValid(email string) bool {
	return strings.Contains(email, "@")
}

func isPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasNumber := false
	hasLetter := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasNumber = true
		} else if unicode.IsLetter(char) {
			hasLetter = true
		}
	}
	if !hasNumber || !hasLetter {
		return false
	}

	return true
}

func isEmailTaken(db *sql.DB, email string) bool {
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email)
	var count int
	err := row.Scan(&count)
	if err != nil || count > 0 {
		return true
	}

	return false
}
