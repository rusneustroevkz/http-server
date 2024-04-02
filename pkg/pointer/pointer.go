package pointer

func Of[t any](arg t) *t {
	return &arg
}
