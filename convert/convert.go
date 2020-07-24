package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

// StringToInt ...
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// StringToInt64 ...
func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// IntToString ...
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// Int64ToString ...
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// ToInt64 convert any numeric value to int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

// toString convert any value to string
func AnyToString(value interface{}) (d string, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int:
		d = strconv.Itoa(int(val.Int()))
	case int64:
		d = strconv.FormatInt(val.Int(), 10)
	case int8, int16, int32:
		d = strconv.Itoa(int(val.Int()))
	case uint, uint32, uint64:
		d = strconv.FormatUint(val.Uint(), 10)
	case uint8, uint16:
		d = strconv.Itoa(int(val.Int()))
	case []byte:
		d = string(val.Bytes())
	case string:
		d = val.String()
	case bool:
		if val.Bool() {
			d = "true"
		} else {
			d = "false"
		}
	default:
		err = fmt.Errorf("not `%T`", value)
	}
	return
}
