package usecase

import (
	user "basiccrud/internal/user/entity"
	"basiccrud/internal/user/repository"
	"context"
	"errors"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, user *user.UserRegister) error
	Login(login *user.UserLogin) (*user.UserRegister, error)
}

type UserInteraction struct {
	UserRepository repository.UserRepository
}

func (u *UserInteraction) RegisterUser(ctx context.Context, user *user.UserRegister) error {
	if len(user.Phone) < 10 {
		return errors.New("username must be at least 3 characters long")
	}

	err := u.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserInteraction) Login(login *user.UserLogin) (*user.UserRegister, error) {
	user, err := u.UserRepository.FindUserByUserName(login.Username)
	if err != nil {
		return nil, err
	}
	if login.Password != user.Password {
		return nil, errors.New("password error")
	}
	return user, nil
}

func UseCase(repo repository.UserRepository) UserUseCase {
	return &UserInteraction{
		UserRepository: repo,
	}
}
