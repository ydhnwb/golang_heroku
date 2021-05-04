package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_heroku/config"
	v1 "github.com/ydhnwb/golang_heroku/handler/v1"
	"github.com/ydhnwb/golang_heroku/middleware"
	"github.com/ydhnwb/golang_heroku/repo"
	"github.com/ydhnwb/golang_heroku/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB               = config.SetupDatabaseConnection()
	userRepo       repo.UserRepository    = repo.NewUserRepo(db)
	productRepo    repo.ProductRepository = repo.NewProductRepo(db)
	authService    service.AuthService    = service.NewAuthService(userRepo)
	jwtService     service.JWTService     = service.NewJWTService()
	userService    service.UserService    = service.NewUserService(userRepo)
	productService service.ProductService = service.NewProductService(productRepo)
	authHandler    v1.AuthHandler         = v1.NewAuthHandler(authService, jwtService, userService)
	userHandler    v1.UserHandler         = v1.NewUserHandler(userService, jwtService)
	productHandler v1.ProductHandler      = v1.NewProductHandler(productService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	productRoutes := server.Group("api/product", middleware.AuthorizeJWT(jwtService))
	{
		productRoutes.GET("/", productHandler.All)
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/:id", productHandler.FindOneProductByID)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
	}

	server.Run()
}
