package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"ledgerApp/src/utils/models"
	"strconv"
)

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateData struct {
			UserID string     `json:"id"`
			User   models.User `json:"user"`
		}
		err := json.NewDecoder(r.Body).Decode(&updateData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Failed to decode update data: %v", err)
			return
		}

		userID, err := strconv.Atoi(updateData.UserID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid user ID")
			return
		}

		if updateData.User.Username == "" || updateData.User.Email == "" || updateData.User.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Missing required fields")
			return
		}

		existingUserID, err := getUserIDByEmail(db, updateData.User.Email)
		if err != nil && err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to retrieve user: %v", err)
			return
		}
		if existingUserID != 0 && existingUserID != userID {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Email is already in use")
			return
		}

		updateData.User.ID = userID
		err = updateUser(db, &updateData.User)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to update user: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User updated successfully")
	}
}

func getUserIDByEmail(db *sql.DB, email string) (int, error) {
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE email = $1", email).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func updateUser(db *sql.DB, user *models.User) error {
	_, err := db.Exec("UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4",
		user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}

	return nil
}
