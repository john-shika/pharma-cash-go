package sqlx

import (
	"nokowebapi/nokocore"
	"strings"
)

type TableNameableImpl interface {
	TableName() string
}

func GetTableName(obj any) string {
	var ok bool
	var nameable TableNameableImpl
	nokocore.KeepVoid(ok, nameable)

	if nokocore.IsNone(obj) {
		return "<nil>"
	}

	// try cast nameable and call method
	if nameable, ok = obj.(TableNameableImpl); ok {
		return nameable.TableName()
	}

	return GetTableNameTypeReflect(obj)
}

func GetTableNameTypeReflect(value any) string {
	var ok bool
	nokocore.KeepVoid(ok)

	val := nokocore.PassValueIndirectReflect(value)
	if !val.IsValid() {
		panic("invalid value")
	}
	method := val.MethodByName("TableName")
	if method.IsValid() {
		results := method.Call(nil)
		if len(results) != 1 || !results[0].IsValid() {
			panic("invalid results")
		}
		name := nokocore.Unwrap(nokocore.CastString(results[0].Interface()))
		return name
	}

	name := nokocore.ToSnakeCase(val.Type().Name())
	return pluralize(name)
}

func pluralize(name string) string {
	if strings.HasSuffix(name, "y") {
		return name + "ies"
	}
	if strings.HasSuffix(name, "s") {
		return name + "es"
	}
	return name + "s"
}
