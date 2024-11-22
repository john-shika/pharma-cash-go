package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/models"
)

type SessionRepositoryImpl interface {
	BaseRepositoryImpl[models.Session]
}

type SessionRepository struct {
	BaseRepository[models.Session]
}

func NewSessionRepository(DB *gorm.DB) SessionRepositoryImpl {
	return &SessionRepository{
		BaseRepository: NewBaseRepository[models.Session](DB),
	}
}
