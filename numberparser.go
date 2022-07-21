package goo

import "strconv"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func ParseNumber[T Number](s string) (T, error) {
	var v T
	switch any(v).(type) {
	case int:
		i, err := strconv.ParseInt(s, 10, 0)
		return T(i), err
	case int8:
		i, err := strconv.ParseInt(s, 10, 8)
		return T(int8(i)), err
	case int16:
		i, err := strconv.ParseInt(s, 10, 16)
		return T(int16(i)), err
	case int32:
		i, err := strconv.ParseInt(s, 10, 32)
		return T(int32(i)), err
	case int64:
		i, err := strconv.ParseInt(s, 10, 64)
		return T(int64(i)), err
	case uint:
		i, err := strconv.ParseUint(s, 10, 0)
		return T(uint(i)), err
	case uint8:
		i, err := strconv.ParseUint(s, 10, 8)
		return T(uint8(i)), err
	case uint16:
		i, err := strconv.ParseUint(s, 10, 16)
		return T(uint16(i)), err
	case uint32:
		i, err := strconv.ParseUint(s, 10, 32)
		return T(uint32(i)), err
	case uint64:
		i, err := strconv.ParseUint(s, 10, 64)
		return T(i), err
	case float32:
		i, err := strconv.ParseFloat(s, 32)
		return T(float32(i)), err
	case float64:
		i, err := strconv.ParseFloat(s, 64)
		return T(i), err
	}
	return 0, nil
}
