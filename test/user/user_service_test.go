package user_test

import (
	"errors"
	"testing"

	"library-management-system/internal/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) CreateUser(u *user.User) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *mockRepository) GetByEmail(email string) (*user.User, error) {
	args := m.Called(email)
	if usr, ok := args.Get(0).(*user.User); ok {
		return usr, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockRepository) GetByID(id uint) (*user.User, error) {
	args := m.Called(id)
	if usr, ok := args.Get(0).(*user.User); ok {
		return usr, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestRegisterUser_Success(t *testing.T) {
	mockRepo := new(mockRepository)
	svc := user.NewService(mockRepo)

	mockRepo.On("CreateUser", mock.AnythingOfType("*user.User")).Return(nil)

	u, err := svc.RegisterUser("John Doe", "john@example.com", "password", "member")
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", u.Name)
	assert.Equal(t, "john@example.com", u.Email)
	assert.NotEqual(t, "password", u.Password)

	mockRepo.AssertExpectations(t)
}

func TestRegisterUser_CreateUserFails(t *testing.T) {
	mockRepo := new(mockRepository)
	svc := user.NewService(mockRepo)

	expectedErr := errors.New("failed to create user")
	mockRepo.On("CreateUser", mock.Anything).Return(expectedErr)

	u, err := svc.RegisterUser("Jane Doe", "jane@example.com", "password", "member")
	assert.Nil(t, u)
	assert.EqualError(t, err, expectedErr.Error())

	mockRepo.AssertExpectations(t)
}

func TestLoginUser_Success(t *testing.T) {
	mockRepo := new(mockRepository)
	svc := user.NewService(mockRepo)

	hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	expectedUser := &user.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: string(hashed),
		Role:     "member",
	}
	mockRepo.On("GetByEmail", "john@example.com").Return(expectedUser, nil)

	u, err := svc.LoginUser("john@example.com", "password")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, u.ID)

	mockRepo.AssertExpectations(t)
}

func TestLoginUser_InvalidCredentials(t *testing.T) {
	mockRepo := new(mockRepository)
	svc := user.NewService(mockRepo)

	hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	expectedUser := &user.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: string(hashed),
		Role:     "member",
	}
	mockRepo.On("GetByEmail", "john@example.com").Return(expectedUser, nil)

	u, err := svc.LoginUser("john@example.com", "wrongpassword")
	assert.Nil(t, u)
	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}
