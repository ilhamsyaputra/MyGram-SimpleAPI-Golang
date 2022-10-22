package exception

func PanicError(err interface{}) {
	if err != nil {
		panic(err)
	}
}
