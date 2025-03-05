**Library Management System with User Authentication and Book Loan Tracking**

using **Golang, Gin, and PostgreSQL**.

## Setup & Run

```bash
git clone https://github.com/ryabchun/library-management-system.git
cd library-management-system
docker compose up --build
```

## Testing

```bash
go test ./test/...
```

## API Endpoints

- **User Service (8081)**: `/register`, `/login`, `/api/profile`
- **Book Service (8082)**: `/books`, `/books/:id/loan`, `/books/:id/return`

**Postman Collection**: See `postman/LibraryManagement.postman_collection.json`

---

**Task description below.**

### Create a Library Management System comprised of two main services (or modules built as separate microservices)

* User Management & Authentication Service
    1. Responsible for handling user registration, login, and secure access via JSON Web Tokens (JWT)
    2. Stores user data, such as name, email, password hash, and role (e.g., “member” or “librarian”)
* Book Management & Loan Service
    1. Manages the library’s inventory (CRUD operations for books)
    2. Tracks the detailed status of each book (available, loaned, missing) and records which user has loaned a book along with the loan dates
    3. Provides endpoints for borrowing (loaning) and returning books, recording the user ID involved as well as timestamps for when a book was loaned (LoanedAt), its due date (DueDate), and when it was returned (ReturnedAt)
* The system uses a relational database (SQLite, Postgres, etc.) via GORM to persist all data
* REST endpoints are developed using a web framework like Gin. Comprehensive logging, error handling, and unit tests should be included for a production-ready system

### Detailed Requirements

**Language & Framework:**

* Golang with Gin (or any preferred REST framework)
* GORM for ORM/database interactions
* Authentication
    1. Implement endpoints for user registration (hash passwords before storage) and login
    2. Generate and return JWT tokens for authenticating subsequent requests
    3. Use middleware to protect sensitive endpoints, ensuring only authenticated users (or users with the right roles) can access loan or book modification endpoints
* Data Persistence:
    1. Persist users, books, and loan records in a relational database
    2. Ensure the Book model includes fields to track the borrower and status changes