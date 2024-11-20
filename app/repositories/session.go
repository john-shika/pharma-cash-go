package repositories

import (
	"gorm.io/gorm"
	"nokowebapi/apis/models"
	"nokowebapi/apis/repositories"
)

type SessionRepositoryImpl interface {
	repositories.BaseRepositoryImpl[models.Session]
}

type SessionRepository struct {
	repositories.BaseRepository[models.Session]
}

func NewSessionRepository(DB *gorm.DB) SessionRepositoryImpl {
	return &SessionRepository{
		BaseRepository: repositories.NewBaseRepository[models.Session](DB),
	}
}
