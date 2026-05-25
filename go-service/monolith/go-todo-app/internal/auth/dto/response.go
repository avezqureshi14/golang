package auth

type AuthResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}