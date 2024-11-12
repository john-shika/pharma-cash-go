package models

import "database/sql"

type User struct {
	Model
	UUID     string         `db:"uuid" gorm:"unique;not null;index" mapstructure:"uuid" json:"uuid" yaml:"uuid"`
	Username string         `db:"username" gorm:"unique;not null;index" mapstructure:"username" json:"username" yaml:"username"`
	Password string         `db:"password" gorm:"not null" mapstructure:"password" json:"password" yaml:"password"`
	Email    sql.NullString `db:"email" gorm:"unique;index" mapstructure:"email" json:"email,omitempty" yaml:"email"`
	Phone    sql.NullString `db:"phone" gorm:"unique;index" mapstructure:"phone" json:"phone,omitempty" yaml:"phone"`
	Admin    string         `db:"admin" gorm:"not null;index" mapstructure:"admin" json:"admin" yaml:"admin"`
	Role     string         `db:"role" gorm:"not null;index" mapstructure:"role" json:"role" yaml:"role"`
	Level    int            `db:"level" gorm:"not null;index" mapstructure:"level" json:"level" yaml:"level"`
}

func (User) TableName() string {
	return "users"
}
