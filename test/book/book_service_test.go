package book_test

import (
	"testing"
	"time"

	"library-management-system/internal/book"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) CreateBook(b *book.Book) error {
	args := m.Called(b)
	return args.Error(0)
}

func (m *mockRepo) GetBook(id uint) (*book.Book, error) {
	args := m.Called(id)
	if bk, ok := args.Get(0).(*book.Book); ok {
		return bk, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockRepo) UpdateBook(b *book.Book) error {
	args := m.Called(b)
	return args.Error(0)
}

func (m *mockRepo) DeleteBook(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockRepo) ListBooks() ([]book.Book, error) {
	args := m.Called()
	if bks, ok := args.Get(0).([]book.Book); ok {
		return bks, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockRepo) CreateLoanRecord(lr *book.LoanRecord) error {
	args := m.Called(lr)
	return args.Error(0)
}

func (m *mockRepo) UpdateLoanRecord(lr *book.LoanRecord) error {
	args := m.Called(lr)
	return args.Error(0)
}

func (m *mockRepo) GetLoanRecord(id uint) (*book.LoanRecord, error) {
	args := m.Called(id)
	if lr, ok := args.Get(0).(*book.LoanRecord); ok {
		return lr, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCreateBook_Success(t *testing.T) {
	repo := new(mockRepo)
	svc := book.NewService(repo)

	repo.On("CreateBook", mock.AnythingOfType("*book.Book")).Return(nil)

	b, err := svc.CreateBook("Go Programming", "Jane Doe", "Introduction to Go")
	assert.NoError(t, err)
	assert.Equal(t, "Go Programming", b.Title)
	assert.Equal(t, "available", b.Status)

	repo.AssertExpectations(t)
}

func TestLoanBook_BookNotAvailable(t *testing.T) {
	repo := new(mockRepo)
	svc := book.NewService(repo)

	repo.On("GetBook", uint(1)).Return(&book.Book{ID: 1, Status: "loaned"}, nil)

	_, err := svc.LoanBook(1, 1, time.Now().Add(7*24*time.Hour))
	assert.Error(t, err)
	assert.Equal(t, "book not available", err.Error())

	repo.AssertExpectations(t)
}

func TestLoanBook_Success(t *testing.T) {
	repo := new(mockRepo)
	svc := book.NewService(repo)

	bk := &book.Book{ID: 1, Status: "available"}
	repo.On("GetBook", uint(1)).Return(bk, nil)
	repo.On("CreateLoanRecord", mock.AnythingOfType("*book.LoanRecord")).Return(nil)
	repo.On("UpdateBook", bk).Return(nil)

	lr, err := svc.LoanBook(1, 1, time.Now().Add(7*24*time.Hour))
	assert.NoError(t, err)
	assert.NotNil(t, lr)
	assert.Equal(t, "loaned", bk.Status)

	repo.AssertExpectations(t)
}

func TestReturnBook_BookNotLoaned(t *testing.T) {
	repo := new(mockRepo)
	svc := book.NewService(repo)

	bk := &book.Book{ID: 1, Status: "available"}
	repo.On("GetBook", uint(1)).Return(bk, nil)

	_, err := svc.ReturnBook(1, 1)
	assert.Error(t, err)
	assert.Equal(t, "book is not loaned", err.Error())

	repo.AssertExpectations(t)
}
