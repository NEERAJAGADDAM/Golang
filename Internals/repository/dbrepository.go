package repository

import "cms/Internals/models"

type DbRepository interface {
	GetAllUsers() ([]*models.User, error)
	GetUserByUserName(userName string) (*models.User, error)
	CreateUser(user models.User) string
}
