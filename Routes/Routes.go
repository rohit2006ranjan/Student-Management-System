package Routes

import (
	"Student-Management-System/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/student-api")
	{
		grp1.GET("student", Controllers.GetStudents)
		grp1.POST("student", Controllers.CreateStudent)
		grp1.GET("student/:student_id", Controllers.GetStudentByID)
		grp1.PUT("student/:student_id", Controllers.UpdateStudent)
		grp1.DELETE("student/:student_id", Controllers.DeleteStudent)
	}
	return r
}
