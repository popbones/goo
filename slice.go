package goo

func MapSlice[E any, V any](s []E, f func(E) V) []V {
	v := make([]V, len(s))
	for i, e := range s {
		v[i] = f(e)
	}
	return v
}

func ReduceSlice[E any, V any](s []E, initial V, f func(V, E) V) V {
	v := initial
	for _, e := range s {
		v = f(v, e)
	}
	return v
}
