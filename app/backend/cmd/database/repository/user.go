package repository

import (
	"fmt"
	"miraculous/cmd/domain"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *domain.UserPayload) (*domain.User, error)
	UpdateUser(userID string) error // upd by id
	CreateUserAttributes(user *domain.UserAttributesPayload) (*domain.UserAttributes, error)
	UpdateUserAttributes(userID string) error
	CreateUserHistory(userHistory *domain.UserHistoryPayload) (*domain.UserHistory, error)
	DeleteUserHistory(ActualProcess string) error // del by process
}

type UserRepository struct {
	db *gorm.DB
}

// CreateUser implements IUserRepository.
func (u *UserRepository) CreateUser(user *domain.UserPayload) (*domain.User, error) {
	var returnedUser *domain.User
	userPayload := &domain.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	err := u.db.Create(&userPayload).Scan(&returnedUser)
	if err.Error != nil {
		return nil, fmt.Errorf("erro ao inserir usuário no banco de dados")
	}

	return returnedUser, nil
}

// CreateUserAttributes implements IUserRepository.
func (u *UserRepository) CreateUserAttributes(user *domain.UserAttributesPayload) (*domain.UserAttributes, error) {
	panic("unimplemented")
}

// CreateUserHistory implements IUserRepository.
func (u *UserRepository) CreateUserHistory(userHistory *domain.UserHistoryPayload) (*domain.UserHistory, error) {
	panic("unimplemented")
}

// DeleteUserHistory implements IUserRepository.
func (u *UserRepository) DeleteUserHistory(ActualProcess string) error {
	panic("unimplemented")
}

// UpdateUser implements IUserRepository.
func (u *UserRepository) UpdateUser(userID string) error {
	panic("unimplemented")
}

// UpdateUserAttributes implements IUserRepository.
func (u *UserRepository) UpdateUserAttributes(userID string) error {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
