package urlvalues

import (
	"reflect"

	"github.com/go-pg/pg/internal/struct_filter"
)

type afterDecodeURLValues interface {
	AfterDecodeURLValues(values Values) error
}

func Decode(value afterDecodeURLValues, values Values) error {
	strct := reflect.Indirect(reflect.ValueOf(value))
	meta := struct_filter.GetStruct(strct.Type())

	for name, values := range values {
		field := meta.Field(name)
		if field != nil && !field.NoDecode() {
			err := field.Scan(field.Value(strct), values)
			if err != nil {
				return err
			}
		}
	}

	if value, ok := value.(afterDecodeURLValues); ok {
		return value.AfterDecodeURLValues(values)
	}

	return nil
}
