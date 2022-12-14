package app

import (
	"github.com/gin-gonic/gin"
	"inventory-api/pkg/app/middleware"
	"net/http"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "weight tracker API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) Routes() *gin.Engine {
	router := s.router
	router.Use(middleware.CORS())

	v1 := router.Group("/v1/api")
	{
		v1.GET("/status", s.ApiStatus())
		v1.POST("/login", s.Login())
		role := v1.Group("/role")
		{
			//role.Use(middleware.JwtTokenCheck)
			role.GET("/list", s.ListRole())
			role.GET("/detail/:id", s.roleDetail())
			role.POST("/create", s.CreateRole())
			role.PUT("/update/:id", s.UpdateRole())
			role.DELETE("/delete/:id", s.DeleteRole())
			role.POST("/batch_delete", s.BatchDeleteRole())
		}
		user := v1.Group("/user")
		{
			user.POST("/create", s.CreateUser())
			user.GET("/list", s.ListUser())
			user.GET("/detail/:id", s.UserDetail())
			user.PUT("/update/:id", s.UpdateUser())
			user.DELETE("/delete/:id", s.DeleteUser())
		}
		productCategory := v1.Group("/product_category")
		{
			//role.Use(middleware.JwtTokenCheck)
			productCategory.GET("/list", s.ListProductCategory())
			productCategory.GET("/detail/:id", s.productCategoryDetail())
			productCategory.POST("/create", s.CreateProductCategory())
			productCategory.PUT("/update/:id", s.UpdateProductCategory())
			productCategory.DELETE("/delete/:id", s.DeleteProductCategory())
			productCategory.POST("/batch_delete", s.BatchDeleteProductCategory())
		}
		product := v1.Group("/product")
		{
			//role.Use(middleware.JwtTokenCheck)
			product.GET("/list/:page/:per_page", s.ListProduct())
			product.GET("/list_by_category/:categoryID", s.ListProductByCategory())
			product.GET("/detail/:id", s.productDetail())
			product.POST("/create", s.CreateProduct())
			product.PUT("/update/:id", s.UpdateProduct())
			product.PUT("/update_category/:productID", s.UpdateCategory())
			product.DELETE("/delete/:id", s.DeleteProduct())
		}
	}

	return router
}
