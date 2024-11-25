package repositories

import (
	"errors"
	"gorm.io/gorm"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
)

type UserRepositoryImpl interface {
	BaseRepositoryImpl[models.User]
	SafeLogin(username string, password string) (*models.User, error)
	Login(username string, password string) (*models.User, error)
}

type UserRepository struct {
	BaseRepository[models.User]
}

func NewUserRepository(DB *gorm.DB) UserRepositoryImpl {
	return &UserRepository{
		BaseRepository: NewBaseRepository[models.User](DB),
	}
}

func (u *UserRepository) SafeLogin(username string, password string) (*models.User, error) {
	var err error
	var user *models.User
	nokocore.KeepVoid(err, user)

	// maybe username is an email or phone
	email := username
	phone := username

	// set preloads
	preloads := []string{"Roles"}

	// try to find the user
	switch {
	case nokocore.EmailRegex().MatchString(email):
		if user, err = u.SafePreFirst(preloads, "email = ?", email); err != nil {
			return nil, err
		}
		break

	case nokocore.PhoneRegex().MatchString(phone):
		if user, err = u.SafePreFirst(preloads, "phone = ?", phone); err != nil {
			return nil, err
		}
		break

	default:
		if user, err = u.SafePreFirst(preloads, "username = ?", username); err != nil {
			return nil, err
		}
		break
	}

	if user != nil {
		pass := nokocore.NewPassword(password)
		if !pass.Equals(user.Password) {
			return nil, errors.New("invalid password")
		}

		return user, nil
	}

	return nil, errors.New("user not found")
}

func (u *UserRepository) Login(username string, password string) (*models.User, error) {
	var err error
	var user *models.User
	nokocore.KeepVoid(err, user)

	// maybe username is an email or phone
	email := username
	phone := username

	// set preloads
	preloads := []string{"Roles"}

	// try to find the user
	switch {
	case nokocore.EmailRegex().MatchString(email):
		if user, err = u.PreFirst(preloads, "email = ?", email); err != nil {
			return nil, err
		}
		break

	case nokocore.PhoneRegex().MatchString(phone):
		if user, err = u.PreFirst(preloads, "phone = ?", phone); err != nil {
			return nil, err
		}
		break

	default:
		if user, err = u.PreFirst(preloads, "username = ?", username); err != nil {
			return nil, err
		}
		break
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
