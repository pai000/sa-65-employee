package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/controller"
	"github.com/pai000/sa-65-project/entity"
	"github.com/pai000/sa-65-project/middlewares"
)

//const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			//user routes
			router.GET("/genders", controller.ListGenders)
			router.GET("/gender/:id", controller.GetGender)
			router.POST("/genders", controller.CreateGender)
			router.PATCH("/genders", controller.UpdateGender)
			router.DELETE("/genders/:id", controller.DeleteGender)

			//memberClass routes
			router.GET("/job_positions", controller.ListJob_Position)
			router.GET("/job_position/:id", controller.GetJob_Position)
			router.POST("/job_positions", controller.CreateJob_Position)
			router.PATCH("/job_positions", controller.UpdateJob_Position)
			router.DELETE("/job_positions/:id", controller.DeleteJob_Position)

			//province routes
			router.GET("/provinces", controller.ListProvince)
			router.GET("/province/:id", controller.GetProvince)
			router.POST("/provinces", controller.CreateProvince)
			router.PATCH("/provinces", controller.UpdateProvince)
			router.DELETE("/provinces/:id", controller.DeleteProvince)

			//role routes
			router.GET("/employees", controller.ListEmployee)
			router.GET("/employee/:id", controller.GetEmployee)
			router.POST("/employees", controller.CreateEmployee)
			router.PATCH("/employees", controller.UpdateEmployee)
			router.DELETE("/employees/:id", controller.DeleteEmployee)

		}
	}

	//Signup User Route
	r.POST("/signup", controller.CreateLoginEmployee)
	// login User Route
	r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("0.0.0.0:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
