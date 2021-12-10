package route

import (
	"mini-ecommerce/internal/handler"
	"mini-ecommerce/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RunAPI ->route setup
func RunAPI(address string) error {

	userHandler := handler.NewUserHandler()
	productHandler := handler.NewProductHandler()
	orderHandler := handler.NewOrderHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Our Mini Ecommerce")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")

	{
		userRoutes.POST("/register", userHandler.AddUser)       // api/user/register
		userRoutes.POST("/signin", userHandler.SignInUser)		// api/user/signin
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())   // api/users
	{
		userProtectedRoutes.GET("/", userHandler.GetAllUser)    // truyen vao path, gin.HandlerFunc
		userProtectedRoutes.GET("/:user", userHandler.GetUser)
		userProtectedRoutes.GET("/:user/products", userHandler.GetProductOrdered)
		userProtectedRoutes.PUT("/:user", userHandler.UpdateUser)
		userProtectedRoutes.DELETE("/:user", userHandler.DeleteUser)
	}

	productRoutes := apiRoutes.Group("/products", middleware.AuthorizeJWT())
	{
		productRoutes.GET("/", productHandler.GetAllProduct)
		productRoutes.GET("/:product", productHandler.GetProduct)
		productRoutes.POST("/", productHandler.AddProduct)
		productRoutes.PUT("/:product", productHandler.UpdateProduct)
		productRoutes.DELETE("/:product", productHandler.DeleteProduct)
	}

	orderRoutes := apiRoutes.Group("/order", middleware.AuthorizeJWT())
	{
		orderRoutes.POST("/product/:product/quantity/:quantity", orderHandler.OrderProduct)
	}

	fileRoutes := r.Group("/file")
	{
		fileRoutes.POST("/single", handler.SingleFile)
		fileRoutes.POST("/multi", handler.MultipleFile)
	}

	return r.Run(address)

}
