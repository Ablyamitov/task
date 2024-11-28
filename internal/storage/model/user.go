package model

type User struct {
	ID        int    `db:"id"`
	LastName  string `db:"last_name"`
	FirstName string `db:"first_name"`
	Gender    string `db:"gender"`
	BirthDate string `db:"birth_date"`
	Phone     string `db:"phone"`
	Role      string `db:"role"`
}
