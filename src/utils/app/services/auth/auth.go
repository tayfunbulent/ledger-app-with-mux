package auth

import (
	"database/sql"
	"fmt"
	"net/http"
)

func AuthenticateUser(next http.HandlerFunc, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.Header.Get("email")
		password := r.Header.Get("password")

		if email == "" || password == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Email and password are required")
			return
		}

		authenticated, err := authenticateUser(db, email, password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Authentication failed: %v", err)
			return
		}

		if !authenticated {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Authentication failed: invalid credentials")
			return
		}

		next(w, r)
	}
}

func authenticateUser(db *sql.DB, email, password string) (bool, error) {
	row := db.QueryRow("SELECT role FROM users WHERE email = $1 AND password = $2", email, password)

	var role string
	err := row.Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("authentication failed: invalid credentials")
		}
		return false, err
	}

	if role == "admin" {
		return true, nil
	}

	return false, nil
}
