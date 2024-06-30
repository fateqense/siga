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

type HistoryResponse struct {
	Cod               string  `json:"cod"`
	DisiplineName     string  `json:"disciplineName"`
	Description       string  `json:"description"`
	FinalGrade        float64 `json:"finalGrade"`
	TotalAbsences     float64 `json:"totalAbsences"`
	PresenceFrequency float64 `json:"presenceFrequency"`
	RenunciationAt    string  `json:"renunciationAt"`
	IsApproved        bool    `json:"isApproved"`
}

type Discipline struct {
	Cod                  string  `json:"cod"`
	Name                 string  `json:"name"`
	TeacherName          string  `json:"teacherName"`
	Class                string  `json:"class"`
	Workload             int     `json:"workload"`
	TotalAbsencesAllowed float64 `json:"totalAbsencesAllowed"`
}
type ScheduleResponse struct {
	Cod        string     `json:"cod"`
	StartsAt   string     `json:"startsAt"`
	EndsAt     string     `json:"endsAt"`
	Discipline Discipline `json:"discipline"`
}
