package responses

type LoginResponse struct {
	Token string `json:"token"`
}

type ProfileResponse struct {
	Ra                 string `json:"ra"`
	Name               string `json:"name"`
	PhotoURL           string `json:"photoUrl"`
	Email              string `json:"email"`
	InstitutionalEmail string `json:"institutionalEmail"`
	AverageGrade       int    `json:"averageGrade"`
	Progression        int    `json:"progression"`
}
