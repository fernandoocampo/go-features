package comparators

type User struct {
	Name     string
	UserName string
	Age      int
}

func Equal[T comparable](x, y T) bool {
	return x == y
}
