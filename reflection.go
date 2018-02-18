package bench

import (
	"reflect"
	"strconv"
)

func (c *FloatCol) AppendReflectionTypeAssertion(row interface{}) error {
	rv := reflect.ValueOf(row)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr:
		c.rows = append(c.rows, float64(rv.Int()))
	case reflect.Float32, reflect.Float64:
		c.rows = append(c.rows, rv.Float())
	case reflect.String:
		num, err := strconv.ParseFloat(rv.String(), 64)
		if err != nil {
			return err
		}
		c.rows = append(c.rows, num)
	default:
		num, err := strconv.ParseFloat(rv.String(), 64)
		if err != nil {
			return err
		}
		c.rows = append(c.rows, num)
	}
	return nil
}
