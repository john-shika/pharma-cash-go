package factories

import (
	"fmt"
	"gorm.io/gorm"
	"nokowebapi/apis/repositories"
	"nokowebapi/console"
	"nokowebapi/nokocore"
)

func BaseFactory[T any](DB *gorm.DB, dummies []T, query string, cb func(dummy T) []any) []T {
	var err error
	var check *T
	nokocore.KeepVoid(err, check)

	repository := repositories.NewBaseRepository[T](DB)
	for i, dummy := range dummies {
		nokocore.KeepVoid(i)

		tableNameType := nokocore.ToSnakeCase(nokocore.GetNameType(dummy))
		args := cb(dummy)

		if cb != nil {
			if check, err = repository.First(query, args...); err != nil {
				console.Warn(err.Error())
				continue
			}

			if check != nil {
				console.Warn(fmt.Sprintf("dummy '%s' already exists. (index=%d)", tableNameType, i))
				continue
			}

			if err = repository.Create(&dummy); err != nil {
				console.Warn(err.Error())
				continue
			}

			console.Warn(fmt.Sprintf("dummy '%s' has been created. (index=%d)", tableNameType, i))
			dummies[i] = dummy
			continue
		}

		console.Warn(fmt.Sprintf("dummy '%s' has been skipped. (index=%d)", tableNameType, i))
	}

	return dummies
}
