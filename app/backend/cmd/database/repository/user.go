package repository

import (
	"miraculous/cmd/domain"
)

type IUserRepository interface {
	CreateUser(user *domain.UserPayload) (*domain.User, error)
	UpdateUser(userID string) error // upd by id
	CreateUserAttributes(user *domain.UserAttributesPayload) (*domain.UserAttributes, error)
	UpdateUserAttributes(userID string) error
	CreateUserHistory(userHistory *domain.UserHistoryPayload) (*domain.UserHistory, error)
	DeleteUserHistory(ActualProcess string) error // del by process
}
