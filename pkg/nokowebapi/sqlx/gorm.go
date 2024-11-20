package sqlx

import (
	"gorm.io/gorm"
	"time"
)

func NewDeletedAt(value time.Time) gorm.DeletedAt {
	return gorm.DeletedAt{
		Time:  value,
		Valid: true,
	}
}
