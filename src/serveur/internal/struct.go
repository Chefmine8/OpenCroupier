package internal

import "github.com/golang-jwt/jwt/v5"

// API
type Response struct {
	Status  string
	Message string
}

// Login
type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
type LoginSQL struct {
	Password string
	Uid      string
}
type LoginResponse struct {
	Token string `json:"token"`
}

// Create
type CreateCustomerRequest struct {
	Name string `json:"name"`
	UID  string `json:"uid"`
}
type CreateUserRequest struct {
	UID      string `json:"uid"`
	UserName string `json:"userName"`
	Pass     string `json:"pass"`
}

// JWT
type Claims struct {
	UserUID  string `json:"uid"`
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}
