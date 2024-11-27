package dto

type SavedResponse struct {
	Data   *Saved    `json:"data"`
	Status bool      `json:"status"`
	Errors *[]string `json:"errors"`
}

type Saved struct {
	Status bool
}
