package service

import (
	"inn/internal/user/model"
	"inn/internal/user/repository"
	"inn/pkg/e"
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) GetUserById(uid uint64) (*model.User, error) {
	return us.repo.FindById(uid)
}

//Exist Determine if the user exists
func (us *UserService) Exist(email string) bool {
	user, _ := us.repo.FindByEmail(email)
	return user != nil
}

func (us *UserService) Login(userLogin *model.UserLoginParam) (*model.User, error) {
	user, err := us.repo.FindByEmail(userLogin.Email)
	if err == nil && user.Pwd != userLogin.Pwd {
		return nil, e.ERROR_WRONGPASSWORD
	}
	return user, err
}

func (us *UserService) Register(userRegister *model.UserRegisterParam) error {
	if us.Exist(userRegister.Email) {
		return e.ERROR_EXIST
	}

	return us.repo.Create(userRegister.ToUser())
}

func (us *UserService) GetAllUsersExcept(uid uint64) ([]*model.User, error) {
	users, err := us.repo.FindUsersByUidIsNot(uid)
	if err != nil {
		return nil, e.ERROR_DBERROR
	}

	return users, err
}
