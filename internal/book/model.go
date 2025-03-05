package book

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Status      string `json:"status"` // "available", "loaned", "missing"
}

type LoanRecord struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	BookID     uint       `json:"book_id"`
	UserID     uint       `json:"user_id"`
	LoanedAt   time.Time  `json:"loaned_at"`
	DueDate    time.Time  `json:"due_date"`
	ReturnedAt *time.Time `json:"returned_at"`
}
