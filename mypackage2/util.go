package mypackage2

import (
	"runtime"
)

type MyType struct {
	Name string
}

func MyFun() MyType {
	return MyType{
		Name: getCallerName(2),
	}
}

func getCallerName(depth int) string {
	_, filename, _, ok := runtime.Caller(depth)
	if !ok {
		panic("Could not get caller name")
	}
	func() {}()
	return filename
}
