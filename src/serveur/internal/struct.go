package internal

// API
type Response struct {
	Status  string
	Message string
}

// SQL
type Utilisateur struct {
	ID    int64
	Nom   string
	Email string
}
