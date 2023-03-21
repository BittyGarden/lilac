package err_tool

// PanicIfNotNil 当错误不为空的时候 panic
func PanicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}
