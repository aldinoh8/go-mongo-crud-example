package dto

type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
