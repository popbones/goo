package goo

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Fallback[T comparable](v T, fallback T) T {
	var newV T
	if newV == v {
		return fallback
	}
	return v
}
