package Models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rohit2006ranjan/Student-Management-System/Config"
)

//GetAllStudents Fetch all student data
func GetAllStudents(student *[]Student) (err error) {
	if err = Config.DB.Find(student).Error; err != nil {
		return err
	}
	return nil
}

//CreateStudent ... Insert New Student data
func CreateStudent(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

//GetStudentByID ... Fetch only one Student by Id
func GetStudentByID(student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}

//UpdateStudent ... Update Student Details
func UpdateStudent(student *Student, id string) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}

//DeleteStudent ... Delete student
func DeleteStudent(student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}
