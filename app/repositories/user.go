package repositories

import (
	"errors"
	"gorm.io/gorm"
	"nokowebapi/apis/models"
	"nokowebapi/apis/repositories"
	"nokowebapi/nokocore"
)

type UserRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models.User]
	SafeLogin(username string, password string) (*models.User, error)
}

type UserRepository struct {
	repositories.BaseRepository[models.User]
}

func NewUserRepository(DB *gorm.DB) UserRepositoryImpl {
	return &UserRepository{
		BaseRepository: repositories.NewBaseRepository[models.User](DB),
	}
}

func (u *UserRepository) SafeLogin(username string, password string) (*models.User, error) {
	var err error
	var user *models.User
	nokocore.KeepVoid(err, user)

	if email := username; nokocore.EmailRegex().MatchString(email) {
		if user, err = u.SafeFirst("email = ?", email); err != nil {
			return nil, err
		}

	} else if phone := username; nokocore.E164Regex().MatchString(phone) {
		if user, err = u.SafeFirst("phone = ?", phone); err != nil {
			return nil, err
		}

	} else {
		if user, err = u.SafeFirst("username = ?", username); err != nil {
			return nil, err
		}
	}

	pass := nokocore.NewPassword(password)
	if user != nil {
		if !pass.Equals(user.Password) {
			return nil, errors.New("invalid password")
		}

		return user, nil
	}

	return nil, nil
}

func (u *UserRepository) Login(username string, password string) (*models.User, error) {
	var err error
	var user *models.User
	nokocore.KeepVoid(err, user)

	if email := username; nokocore.EmailRegex().MatchString(email) {
		if user, err = u.First("email = ?", email); err != nil {
			return nil, err
		}

	} else if phone := username; nokocore.E164Regex().MatchString(phone) {
		if user, err = u.First("phone = ?", phone); err != nil {
			return nil, err
		}

	} else {
		if user, err = u.First("username = ?", username); err != nil {
			return nil, err
		}
	}

	pass := nokocore.NewPassword(password)
	if user != nil {
		if !pass.Equals(user.Password) {
			return nil, errors.New("invalid password")
		}

		return user, nil
	}

	return nil, nil
}
