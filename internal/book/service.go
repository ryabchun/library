package book

import (
	"errors"
	"time"
)

type Service interface {
	CreateBook(title, author, description string) (*Book, error)
	UpdateBook(id uint, title, author, description string) (*Book, error)
	GetBook(id uint) (*Book, error)
	DeleteBook(id uint) error
	ListBooks() ([]Book, error)

	LoanBook(bookID, userID uint, dueDate time.Time) (*LoanRecord, error)
	ReturnBook(bookID, userID uint) (*LoanRecord, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateBook(title, author, description string) (*Book, error) {
	b := &Book{
		Title:       title,
		Author:      author,
		Description: description,
		Status:      "available",
	}
	err := s.repo.CreateBook(b)
	return b, err
}

func (s *service) UpdateBook(id uint, title, author, description string) (*Book, error) {
	b, err := s.repo.GetBook(id)
	if err != nil {
		return nil, err
	}
	b.Title = title
	b.Author = author
	b.Description = description
	err = s.repo.UpdateBook(b)
	return b, err
}

func (s *service) GetBook(id uint) (*Book, error) {
	return s.repo.GetBook(id)
}

func (s *service) DeleteBook(id uint) error {
	return s.repo.DeleteBook(id)
}

func (s *service) ListBooks() ([]Book, error) {
	return s.repo.ListBooks()
}

func (s *service) LoanBook(bookID, userID uint, dueDate time.Time) (*LoanRecord, error) {
	b, err := s.repo.GetBook(bookID)
	if err != nil {
		return nil, err
	}

	if b.Status != "available" {
		return nil, errors.New("book not available")
	}

	lr := &LoanRecord{
		BookID:   bookID,
		UserID:   userID,
		LoanedAt: time.Now(),
		DueDate:  dueDate,
	}
	err = s.repo.CreateLoanRecord(lr)
	if err != nil {
		return nil, err
	}

	b.Status = "loaned"
	if err := s.repo.UpdateBook(b); err != nil {
		return nil, err
	}
	return lr, nil
}

func (s *service) ReturnBook(bookID, userID uint) (*LoanRecord, error) {
	b, err := s.repo.GetBook(bookID)
	if err != nil {
		return nil, err
	}

	if b.Status != "loaned" {
		return nil, errors.New("book is not loaned")
	}

	lr, err := s.repo.GetLoanRecord(b.ID)
	if err != nil {
		return nil, err
	}
	if lr.ReturnedAt != nil {
		return nil, errors.New("already returned")
	}

	now := time.Now()
	lr.ReturnedAt = &now
	err = s.repo.UpdateLoanRecord(lr)
	if err != nil {
		return nil, err
	}

	b.Status = "available"
	err = s.repo.UpdateBook(b)
	return lr, err
}
