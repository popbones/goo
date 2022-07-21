package goo

func MapSlice[E any, V any](s []E, f func(E) V) []V {
	var v []V
	for _, e := range s {
		v = append(v, f(e))
	}
	return v
}

func ReduceSlice[E any, V any](s []E, initial V, f func(V, E) V) V {
	var v V
	for _, e := range s {
		v = f(v, e)
	}
	return v
}
