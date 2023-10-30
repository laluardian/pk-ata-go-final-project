package repositories

import (
	"github.com/google/uuid"
	"github.com/laluardian/pk-ata-go-final-project/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
	FindMany() ([]models.Book, error)
	FindById(userId uuid.UUID) (models.Book, error)
	Update(book *models.Book) error
	Delete(bookId uuid.UUID) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (b *bookRepository) Create(book *models.Book) error {
	return b.db.Create(&book).Error
}

func (b *bookRepository) FindMany() (books []models.Book, err error) {
	err = b.db.Find(&books).Error
	return books, err

}

func (b *bookRepository) FindById(bookId uuid.UUID) (book models.Book, err error) {
	err = b.db.First(&book, "id = ?", bookId).Error
	return book, err
}

func (b *bookRepository) Update(book *models.Book) error {
	return b.db.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(&book).Error
}

func (b *bookRepository) Delete(bookId uuid.UUID) error {
	var book models.Book
	return b.db.Delete(&book, "id = ?", bookId).Error
}
