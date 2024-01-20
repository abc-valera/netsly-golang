package codeerr

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustWithVal[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
