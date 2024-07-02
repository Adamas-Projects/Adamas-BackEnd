package reqs

type InstitutionCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CNPJ     string    `json:"cnpj"`
}
