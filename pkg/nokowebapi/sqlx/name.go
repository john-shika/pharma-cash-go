package sqlx

import (
	"nokowebapi/nokocore"
	"strings"
)

type NameableImpl interface {
	TableName() string
}

func GetTableName(obj any) string {
	var ok bool
	var nameable NameableImpl
	nokocore.KeepVoid(ok, nameable)

	if nokocore.IsNone(obj) {
		return "none"
	}

	// try cast nameable and call method
	if nameable, ok = nokocore.Cast[NameableImpl](obj); ok {
		return nameable.TableName()
	}

	return GetTableNameReflect(obj)
}

func GetTableNameReflect(value any) string {
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

// pluralize method, not perfect but good enough
func pluralize(name string) string {
	if strings.HasSuffix(name, "ies") {
		return name
	}
	if value, ok := strings.CutSuffix(name, "y"); ok {
		return value + "ies"
	}
	if strings.HasSuffix(name, "es") {
		return name
	}
	if strings.HasSuffix(name, "s") {
		return name + "es"
	}
	return name + "s"
}
