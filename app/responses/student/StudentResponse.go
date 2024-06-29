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

type Exam struct {
	Title    string `json:"title"`
	StartsAt string `json:"startsAt"`
	Grade    int    `json:"grade"`
}
type PartialGradeResponse struct {
	Cod           string  `json:"cod"`
	DisiplineName string  `json:"disciplineName"`
	AverageGrade  float64 `json:"averageGrade"`
	Exams         []*Exam
}

type Lesson struct {
	Title     string  `json:"title"`
	Date      string  `json:"date"`
	Presences float64 `json:"presences"`
	Absences  float64 `json:"absences"`
}
type PartialAbsenceRespose struct {
	Cod            string   `json:"cod"`
	DisiplineName  string   `json:"disciplineName"`
	TotalPresences float64  `json:"totalPresences"`
	TotalAbsences  float64  `json:"totalAbsences"`
	Lessons        []Lesson `json:"lessons"`
}
