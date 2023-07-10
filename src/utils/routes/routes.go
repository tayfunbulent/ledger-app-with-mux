package routes

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	"ledgerApp/src/utils/app/services/auth"
	"ledgerApp/src/utils/app/services/transactions"
	"ledgerApp/src/utils/app/services/users"
	"ledgerApp/src/utils/app/services/wallets"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	/* POST requests */
	/* User */
	router.HandleFunc("/app/services/users/create", users.CreateUser(db)).Methods(http.MethodPost)
	router.HandleFunc("/app/services/wallets/update", wallets.UpdateWalletBalance(db)).Methods(http.MethodPost)
	router.HandleFunc("/app/services/transactions/create", transactions.CreateTransaction(db)).Methods(http.MethodPost)

	/* Admin */
	router.HandleFunc("/app/services/users/update-user",
		auth.AuthenticateUser(users.UpdateUser(db), db)).Methods(http.MethodPost)
	router.HandleFunc("/app/services/users/delete-user",
		auth.AuthenticateUser(users.DeleteUser(db), db)).Methods(http.MethodPost)
	router.HandleFunc("/app/services/users/update-user-role",
		auth.AuthenticateUser(users.UpdateUserRole(db), db)).Methods(http.MethodPost)
	router.HandleFunc("/app/services/wallets/update-wallet-balance-by-user-id",
		auth.AuthenticateUser(wallets.UpdateWalletBalanceByUserID(db), db)).Methods(http.MethodPost)
	router.HandleFunc("/app/services/transactions/get-wallet-balance-at-time",
		auth.AuthenticateUser(wallets.GetWalletBalanceAtTime(db), db)).Methods(http.MethodPost)

	/* GET requests */
	/* User */
	router.HandleFunc("/app/services/wallets/get", wallets.GetWallet(db)).Methods(http.MethodGet)
	router.HandleFunc("/app/services/transactions/get", transactions.GetTransactions(db)).Methods(http.MethodGet)

	/* Admin */
	router.HandleFunc("/app/services/users/get-all-users",
		auth.AuthenticateUser(users.GetAllUsers(db), db)).Methods(http.MethodGet)
	router.HandleFunc("/app/services/users/get-all-users-with-wallet",
		auth.AuthenticateUser(users.GetAllUsersWithWallet(db), db)).Methods(http.MethodGet)
	router.HandleFunc("/app/services/wallets/get-all-wallets",
		auth.AuthenticateUser(wallets.GetAllWallets(db), db)).Methods(http.MethodGet)
	router.HandleFunc("/app/services/users/get-user-by-id/{id}",
		auth.AuthenticateUser(users.GetUserByID(db), db)).Methods(http.MethodGet)
	router.HandleFunc("/app/services/wallets/get-wallet-by-id/{id}",
		auth.AuthenticateUser(wallets.GetWalletByID(db), db)).Methods(http.MethodGet)

	return router
}
