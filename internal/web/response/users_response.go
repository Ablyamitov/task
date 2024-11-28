package response

type UsersResponse struct {
	Data   *Fetched  `json:"data"`
	Status bool      `json:"status"`
	Errors *[]string `json:"errors"`
}

type Fetched struct {
	Users []*UserDTO
}
