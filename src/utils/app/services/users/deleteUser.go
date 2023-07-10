package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var deleteData struct {
			UserID string `json:"id"`
		}
		err := json.NewDecoder(r.Body).Decode(&deleteData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Failed to decode delete data: %v", err)
			return
		}

		userID, err := strconv.Atoi(deleteData.UserID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Invalid user ID")
			return
		}

		err = deleteUser(db, userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to delete user: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User deleted successfully")
	}
}

func deleteUser(db *sql.DB, userID int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		return err
	}

	return nil
}
