package internal

// API
type Response struct {
	Status  string
	Message string
}

type CreateCustomerRequest struct {
	Name string `json:"name"`
	UID  string `json:"uid"`
}
type CreateUserRequest struct {
	UID      string `json:"uid"`
	UserName string `json:"userName"`
	Pass     string `json:"pass"`
}
