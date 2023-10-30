package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/laluardian/pk-ata-go-final-project/controllers"
	"gorm.io/gorm"
)

func InitApi(db *gorm.DB) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	bc := controllers.NewBookController(db)

	e.POST("/", bc.AddBook)
	e.GET("/", bc.GetBooks)
	e.GET("/:id", bc.GetBookById)
	e.PUT("/:id", bc.UpdateBook)
	e.DELETE("/:id", bc.DeleteBook)

	return e
}
