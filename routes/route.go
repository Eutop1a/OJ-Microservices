package routes

import (
	"OnlineJudge/controller"
	_ "OnlineJudge/docs"
	"OnlineJudge/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
)

// SetUp initializes and configures a Gin Engine with routes for your application.
// It sets up various middleware, CORS headers, and defines API routes.
//
// Parameters:
//   - mode: The execution mode of the application (e.g., gin.DebugMode, gin.ReleaseMode).
//
// Returns:
//   - *gin.Engine: The configured Gin Engine.
func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		// Set Gin to release mode
		gin.SetMode(gin.ReleaseMode)
	}
	// Create a new Gin Engine with default middleware
	r := gin.Default()

	// Set up CORS headers middleware
	r.Use(middlewares.Cors())
	// Group for authenticated routes
	p := r.Group("/api/v1")

	{
		// Swagger documentation route
		p.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

		// Public endpoints
		p.POST("/login", controller.Login)
		p.POST("/register", controller.Register)
		p.POST("/check-username", controller.CheckUserName)
		p.POST("/send-email-code", controller.SendEmailCode) // 发送验证码
		p.POST("/send-code", controller.SendCode)            // 发送验证码
	}

	p.Use(middlewares.AuthorizationMiddleWare)
	{
		// Authenticated endpoints
		p.GET("/users/:id", controller.GetUserDetail)
		p.DELETE("/users/:id", controller.DeleteUser)
		p.PUT("/users/:id", controller.UpdateUserDetail)
		p.POST("/sort-users-by-score", controller.SortByScore)
		p.POST("/change-user-score", controller.AddScoreToUser)
		p.POST("/check-token", controller.CheckTokenValidity)
		p.GET("/list-problems", controller.GetProblemsList)
		p.POST("/problem/add", controller.AddProblem)
		p.GET("/problem/:id", controller.GetProblem)

		// p.POST("/admin", controller.RegAdmin)
	} // CRUD

	// Group for file-related routes
	{
		p.GET("/problem/file/:id", controller.GetProblemFiles)
		p.POST("/problem/file/add/:id", controller.AddFilesToProblem)
	} // Files

	// Define a route for code judging
	p.POST("/judge-code", controller.JudgeCode)

	// Ping route for testing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})

	// 404 route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})

	return r
}
