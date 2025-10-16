package struct2env

import (
	"os"
	"reflect"
)

func Bind(cfg interface{}) {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		field := t.Field(i)

		if !field.IsExported() {
			continue
		}
		name, ok := field.Tag.Lookup("env")
		if !ok {
			continue
		}
		switch f.Kind() {
		case reflect.Struct:
			Bind(f.Addr().Interface())
			continue
		case reflect.Pointer:
			if f.Type().Elem().Kind() == reflect.Struct && !f.IsNil() {
				Bind(f.Interface())
				continue
			}
		case reflect.String:
			value, ok := os.LookupEnv(name)
			if ok {
				f.SetString(value)
			}
		}
	}
}
