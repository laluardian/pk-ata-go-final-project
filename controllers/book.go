package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/laluardian/pk-ata-go-final-project/libs"
	"github.com/laluardian/pk-ata-go-final-project/models"
	"github.com/laluardian/pk-ata-go-final-project/repositories"
	"gorm.io/gorm"
)

type BookController interface {
	AddBook(c echo.Context) error
	GetBooks(c echo.Context) error
	GetBookById(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
}

type bookController struct {
	repo repositories.BookRepository
}

func NewBookController(db *gorm.DB) BookController {
	return &bookController{repositories.NewBookRepository(db)}
}

// AddBook implements BookController.
func (b *bookController) AddBook(c echo.Context) error {
	var book models.Book
	c.Bind(&book)

	if err := b.repo.Create(&book); err != nil {
		c.JSON(http.StatusInternalServerError, &libs.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return err
	}

	return c.JSON(http.StatusCreated, &libs.ApiResponse{
		Success: true,
		Message: "a new book is added",
		Data:    book,
	})

}

func (b *bookController) GetBooks(c echo.Context) error {
	books, err := b.repo.FindMany()

	if err != nil {
		c.JSON(http.StatusInternalServerError, &libs.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return err
	}

	return c.JSON(http.StatusOK, &libs.ApiResponse{
		Success: true,
		Message: "the books are retrieved",
		Data:    books,
	})
}

func (b *bookController) GetBookById(c echo.Context) error {
	bookId, _ := uuid.Parse(c.Param("id"))
	book, err := b.repo.FindById(bookId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &libs.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return err
	}

	return c.JSON(http.StatusOK, &libs.ApiResponse{
		Success: true,
		Message: "the book is retrieved",
		Data:    book,
	})
}

func (b *bookController) UpdateBook(c echo.Context) error {
	bookId, _ := uuid.Parse(c.Param("id"))
	book, err := b.repo.FindById(bookId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &libs.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return err
	}

	c.Bind(&book)
	book.ID = bookId

	if err := b.repo.Update(&book); err != nil {
		c.JSON(http.StatusInternalServerError, &libs.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return err
	}

	return c.JSON(http.StatusOK, &libs.ApiResponse{
		Success: true,
		Message: "the book is updated",
		Data:    book,
	})
}

func (b *bookController) DeleteBook(c echo.Context) error {
	bookId, _ := uuid.Parse(c.Param("id"))
	err := b.repo.Delete(bookId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &libs.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return err
	}

	return c.JSON(http.StatusOK, &libs.ApiResponse{
		Success: true,
		Message: "the book is deleted",
	})
}
