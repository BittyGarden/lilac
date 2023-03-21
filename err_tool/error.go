// Package err_tool 错误处理工具包
package err_tool

// PanicIfNoNil 当错误不为空的时候panic
func PanicIfNoNil(err error) {
	if err != nil {
		panic(err)
	}
}
