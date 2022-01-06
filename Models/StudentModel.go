package Models

type Student struct {
	Student_id      uint64 `gorm:"primary_key"`
	First_name      string
	Last_name       string
	Mobile          string `gorm:"unique"`
	Year_of_joining string
	Email_id        string
}

func (b *Student) TableName() string {
	return "student"
}
