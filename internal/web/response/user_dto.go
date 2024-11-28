package response

type UserDTO struct {
	ID        int    `json:"id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	Phone     string `json:"phone"`
}
