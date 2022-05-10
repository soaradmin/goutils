package internal

import "runtime"

//Caller return the caller file and code line
func Caller() (string, int) {
	pc, _, b, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name(), b
}
