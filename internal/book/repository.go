package book

import "gorm.io/gorm"

type Repository interface {
	CreateBook(b *Book) error
	GetBook(id uint) (*Book, error)
	UpdateBook(b *Book) error
	DeleteBook(id uint) error
	ListBooks() ([]Book, error)

	CreateLoanRecord(lr *LoanRecord) error
	UpdateLoanRecord(lr *LoanRecord) error
	GetLoanRecord(id uint) (*LoanRecord, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateBook(b *Book) error {
	return r.db.Create(b).Error
}

func (r *repository) GetBook(id uint) (*Book, error) {
	var book Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *repository) UpdateBook(b *Book) error {
	return r.db.Save(b).Error
}

func (r *repository) DeleteBook(id uint) error {
	return r.db.Delete(&Book{}, id).Error
}

func (r *repository) ListBooks() ([]Book, error) {
	var books []Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *repository) CreateLoanRecord(lr *LoanRecord) error {
	return r.db.Create(lr).Error
}

func (r *repository) UpdateLoanRecord(lr *LoanRecord) error {
	return r.db.Save(lr).Error
}

func (r *repository) GetLoanRecord(id uint) (*LoanRecord, error) {
	var lr LoanRecord
	if err := r.db.First(&lr, id).Error; err != nil {
		return nil, err
	}
	return &lr, nil
}
