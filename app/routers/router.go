package routers

import (
	"clean-arch/app/middlewares"
	"clean-arch/features/user/data"
	_userHandler "clean-arch/features/user/handler"
	_userService "clean-arch/features/user/service"

	_productData "clean-arch/features/product/data"
	_productHandler "clean-arch/features/product/handler"
	_productService "clean-arch/features/product/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	// Inisialisasi data dan service untuk entitas User
	userData := data.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	// Inisialisasi data dan service untuk entitas Product
	productData := _productData.NewProduct(db)
	productService := _productService.NewProductService(productData)
	productHandlerAPI := _productHandler.NewProductHandler(productService)

	// Definisikan rute untuk entitas User
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/users", userHandlerAPI.GetAllUsers, middlewares.JWTMiddleware())
	e.PUT("/users/:user_id", userHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/users/:user_id", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())

	// Definisikan rute untuk entitas Product
	e.POST("/products", productHandlerAPI.CreateProduct)
	e.GET("/products", productHandlerAPI.GetAllProducts, middlewares.JWTMiddleware())
	e.GET("/products/:product_id", productHandlerAPI.GetProductByID, middlewares.JWTMiddleware())
	e.PUT("/products/:product_id", productHandlerAPI.UpdateProduct, middlewares.JWTMiddleware())
	e.DELETE("/products/:product_id", productHandlerAPI.DeleteProduct, middlewares.JWTMiddleware())
	e.GET("/users/:user_id/products", productHandlerAPI.GetProductsByUserID, middlewares.JWTMiddleware())
}
