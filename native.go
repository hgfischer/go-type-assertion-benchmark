package bench

import (
	"fmt"
	"strconv"
)

func (c *FloatCol) AppendNativeTypeAssertion(row interface{}) error {
	switch value := row.(type) {
	case float32:
		c.rows = append(c.rows, float64(value))
	case float64:
		c.rows = append(c.rows, float64(value))
	case int:
		c.rows = append(c.rows, float64(value))
	case int8:
		c.rows = append(c.rows, float64(value))
	case int16:
		c.rows = append(c.rows, float64(value))
	case int32:
		c.rows = append(c.rows, float64(value))
	case int64:
		c.rows = append(c.rows, float64(value))
	case uint:
		c.rows = append(c.rows, float64(value))
	case uint8:
		c.rows = append(c.rows, float64(value))
	case uint16:
		c.rows = append(c.rows, float64(value))
	case uint32:
		c.rows = append(c.rows, float64(value))
	case uint64:
		c.rows = append(c.rows, float64(value))
	default:
		num, err := strconv.ParseFloat(fmt.Sprint(row), 64)
		if err != nil {
			return err
		}
		c.rows = append(c.rows, num)
	}
	return nil
}
