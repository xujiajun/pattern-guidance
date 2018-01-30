package decorator

type obj func(string) string

func Decorate(fn obj) obj {
	return func(base string) string {
		ret := fn(base)
		ret = ret + " and shoes"
		return ret
	}
}

func Dressing(cloth string) string {
	return "dressing " + cloth
}
