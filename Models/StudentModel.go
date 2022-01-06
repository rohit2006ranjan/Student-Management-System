package Models

type Student struct {
	Student_id      uint   `json:"student_id" gorm:primaryKey`
	First_name      string `json:"first_name"`
	Last_name       string `json:"last_name"`
	Mobile          string `json:"mobile" gorm:unique`
	Year_of_joining string `json:"year_of_joining"`
	Email_id        string `json:"email_id"`
}

func (b *Student) TableName() string {
	return "student"
}
