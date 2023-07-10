package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateUserRole(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateData struct {
			UserID string `json:"id"`
			Role   string `json:"role"`
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
			fmt.Fprint(w, "Invalid user ID")
			return
		}

		err = updateUserRole(db, userID, updateData.Role)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to update user role: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User role updated successfully")
	}
}

func updateUserRole(db *sql.DB, userID int, role string) error {
	_, err := db.Exec("UPDATE users SET role = $1 WHERE id = $2", role, userID)
	if err != nil {
		return err
	}

	return nil
}
