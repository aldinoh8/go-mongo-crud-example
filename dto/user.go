package dto

type RegisterBody struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginBody struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type TopupBody struct {
	Amount int `json:"amount"`
}
