package Models

type Student struct {
	student_id      uint   `json:"student_id"`
	first_name      string `json:"firstName"`
	last_name       string `json:"lastName"`
	mobile          string `json:"mobile"`
	year_of_joining string `json:"yearOfJoining"`
	email_id        string `json:"emailId"`
}

func (b *Student) TableName() string {
	return "student"
}
