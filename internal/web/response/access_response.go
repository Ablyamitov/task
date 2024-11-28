package response

type AccessResponse struct {
	Data   *Access   `json:"data"`
	Status bool      `json:"status"`
	Errors *[]string `json:"errors"`
}

type Access struct {
	Status bool
}
