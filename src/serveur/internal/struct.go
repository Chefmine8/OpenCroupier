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

// SQL
type Utilisateur struct {
	ID    int64
	Nom   string
	Email string
}
