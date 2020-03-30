package orm

import (
	"inn/internal/user/model"
	"inn/internal/user/repository"
	"inn/pkg/e"
)

type userRepository struct{}

//NewUserRepository return userRepository obj
func NewUserRepository() repository.IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) Create(user *model.User) error {
	return db.Create(user).Error
}

func (ur *userRepository) FindById(uid uint64) (*model.User, error) {
	res := new(model.User)
	q := db.Where("uid = ?", uid).First(res)
	if q.Error != nil {
		if q.RecordNotFound() {
			return nil, e.ERROR_NOTEXIST
		}
	}

	return res, q.Error
}

func (ur *userRepository) FindByEmail(email string) (*model.User, error) {
	res := new(model.User)
	q := db.Where("email = ?", email).First(res)
	if q.Error != nil {
		if q.RecordNotFound() {
			return nil, e.ERROR_NOTEXIST
		}
	}

	return res, q.Error
}

func (ur *userRepository) FindUsersByUidIsNot(uid uint64) ([]*model.User, error) {
	var res []*model.User
	q := db.Where("uid != ?", uid).Find(&res)
	if q.Error != nil {
		if q.RecordNotFound() {
			return res, nil
		}
	}

	return res, q.Error
}
