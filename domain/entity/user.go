package entity

// User 아래 필드들은 Google OAuth2를 이용해 가져옴 (balance 필드 제외)
type User struct {
	Id      string `db:"id" json:"id"`
	Email   string `db:"email" json:"email"`
	Name    string `db:"name" json:"name"`
	Balance uint64 `db:"balance" json:"balance,omitempty"`
}

func NewUser(id, name, email string, balance uint64) *User {
	return &User{
		id,
		name,
		email,
		balance,
	}

}
