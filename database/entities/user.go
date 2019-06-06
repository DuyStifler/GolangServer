package entities

type User struct {
	ID       int64  `db:"id" json:"id"`
	PublicID string `db:"public_id" json:"public_id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}