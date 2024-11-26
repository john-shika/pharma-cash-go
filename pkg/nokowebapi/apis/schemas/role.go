package schemas

import (
	"github.com/google/uuid"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
)

type RoleBody struct {
	Name string `mapstructure:"name" json:"name" form:"name" validate:"alphanum,min=1"`
}

func ToRoleModel(role *RoleBody) *models.Role {
	if role != nil {
		return &models.Role{
			RoleName: nokocore.ToPascalCase(role.Name),
		}
	}

	return nil
}

type RoleResult struct {
	UUID      uuid.UUID `mapstructure:"uuid" json:"uuid"`
	Name      string    `mapstructure:"name" json:"name"`
	CreatedAt string    `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt string    `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt string    `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
}

func ToRoleResult(role *models.Role) RoleResult {
	if role != nil {
		createdAt := nokocore.ToTimeUtcStringISO8601(role.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(role.UpdatedAt)
		var deletedAt string
		if role.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(role.DeletedAt.Time)
		}
		return RoleResult{
			UUID:      role.UUID,
			Name:      role.RoleName,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		}
	}

	return RoleResult{}
}
