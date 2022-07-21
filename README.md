# goo
goo - Go goodies 

A joke package that include things that should be in the STD since generics.

Most of them are written by Copilot.`

- `func Must[T any](v T, err error) T`

    Returns `v` if `err` is `nil`, otherwise panic.

- `func Fallback[T comparable](v T, fallback T) T`

    Returns `v` if `v` is not empty value for `T`, otherwise return `fallback`.

- `func MapSlice[E any, V any](s []E, f func(E) V) []V`

    Map slice.

- `func ReduceSlice[E any, V any](s []E, initial V, f func(V, E) V) V`

    Reduce slice.

- `func ParseNumber[T Number](s string) (T, error)`

    Parse a number from string.

- `func Ptr[T any](v T) *T`

    Returns pointer to value.
