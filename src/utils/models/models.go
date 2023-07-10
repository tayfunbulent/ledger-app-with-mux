package models

// User model
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Transaction model
type Transaction struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Balance     float64 `json:"balance"`
	CreatedAt   string  `json:"created_at"`
}

// Wallet modeli
type Wallet struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
}

// Kullanıcı ve cüzdan bilgilerini içeren yapı
type UserWithWallet struct {
	User   *User   `json:"user"`
	Wallet *Wallet `json:"wallet"`
}
