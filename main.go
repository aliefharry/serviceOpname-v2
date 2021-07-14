package main

import (
	"serviceOpname-v2/controller"
	"serviceOpname-v2/config"
	"serviceOpname-v2/middleware"
	"serviceOpname-v2/service"
	"serviceOpname-v2/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
// import "net/http"
)

var (
	db						*gorm.DB 						= config.SetUpDB()
	userRepository 			repository.UserRepository 		= repository.NewUserRepository(db)
	opnameRepository 		repository.OpnameRepository 	= repository.NewOpnameRepository(db)
	jwtService 				service.JWTService 				= service.NewJWTService()
	
	authService 			service.AuthService 			= service.NewAuthService(userRepository)
	authController 			controller.AuthController 		= controller.NewAuthController(authService, jwtService)
	
	userService 			service.UserService 			= service.NewUserService(userRepository)
	userController 			controller.UserController 		= controller.NewUserController(userService, jwtService)

	opnameService			service.OpnameService 			= service.NewOpnameService(opnameRepository)
	opnameController 		controller.OpnameController 	= controller.NewOpnameController(opnameService, jwtService)
)

func main() {
	defer config.CloseConnDb(db)
	r := gin.Default()
	
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("profile", userController.Profile)
		userRoutes.PUT("uodateProfile", userController.Update)
	}

	opnameRoutes := r.Group("/")
	{
		opnameRoutes.GET("/dataOpname", opnameController.All)
		opnameRoutes.GET("/paginationOpname", opnameController.Paginates)
		opnameRoutes.GET("/:id", opnameController.FindById)
		opnameRoutes.PUT("/:id", opnameController.Update)
	}
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}