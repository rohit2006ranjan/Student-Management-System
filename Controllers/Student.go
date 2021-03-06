package Controllers

import (
	"fmt"
	"net/http"

	"Student-Management-System/Models"

	"github.com/gin-gonic/gin"
)

//GetStudents ... Get all students
func GetStudents(c *gin.Context) {
	var student []Models.Student
	err := Models.GetAllStudents(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//CreateStudent ... Create Student
func CreateStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	err := Models.CreateStudent(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//GetStudentByID ... Get the student by id
func GetStudentByID(c *gin.Context) {
	student_id := c.Params.ByName("student_id")
	var student Models.Student
	err := Models.GetStudentByID(&student, student_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//UpdateStudent ... Update the student information
func UpdateStudent(c *gin.Context) {
	var student Models.Student
	student_id := c.Params.ByName("student_id")
	err := Models.GetStudentByID(&student, student_id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	c.BindJSON(&student)
	err = Models.UpdateStudent(&student, student_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//DeleteStudent ... Delete the student
func DeleteStudent(c *gin.Context) {
	var student Models.Student
	student_id := c.Params.ByName("student_id")
	err := Models.DeleteStudent(&student, student_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"student_id" + student_id: "is deleted"})
	}
}
