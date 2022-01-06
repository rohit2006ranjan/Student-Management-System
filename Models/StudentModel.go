package Models

type Student1 struct {
	Student_id      uint   `json:"student_id"`
	First_name      string `json:"first_name"`
	Last_name       string `json:"last_name"`
	Mobile          string `json:"mobile"`
	Year_of_joining string `json:"year_of_joining"`
	Email_id        string `json:"email_id"`
}

func (b *Student1) TableName() string {
	return "student"
}
