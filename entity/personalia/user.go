package personalia

type User struct {
	UserID   string `db:"user_id"`
	UserName string `db:"user_name"`
	FullName string `db:"full_name"`
	Password string `db:"password"`
	IsAdmin  bool   `db:"is_admin"`
}
