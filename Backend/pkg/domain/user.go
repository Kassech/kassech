package domain

type User struct {
	ID       uint
	Email    string
	Password string // Hashed password
	Name     string
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}
