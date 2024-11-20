package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/models"
	"nokowebapi/apis/repositories"
)

type SessionRepository struct {
	repositories.BaseRepository[models.Session]
}

func NewSessionRepository(DB *gorm.DB) SessionRepository {
	return SessionRepository{
		BaseRepository: repositories.NewBaseRepository[models.Session](DB),
	}
}

func (u *SessionRepository) Find(wheres ...any) (*models.Session, error) {
	return u.BaseRepository.Find(wheres...)
}

func (u *SessionRepository) Create(model *models.Session) error {
	return u.BaseRepository.Create(model)
}
