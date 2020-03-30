package repository

import (
	"inn/internal/user/model"
)

//mockgen -destination user_mock.go -source user.go -package repository
type IUserRepository interface {
	Create(*model.User) error
	FindById(uid uint64) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindUsersByUidIsNot(uid uint64) ([]*model.User, error)
}
