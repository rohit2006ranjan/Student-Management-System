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
		grp1.GET("student", Controllers.GetUsers)
		grp1.POST("student", Controllers.CreateUser)
		grp1.GET("student/:id", Controllers.GetUserByID)
		grp1.PUT("student/:id", Controllers.UpdateUser)
		grp1.DELETE("student/:id", Controllers.DeleteUser)
	}
	return r
}
