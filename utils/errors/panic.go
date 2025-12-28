package errors

import "fmt"

func PanicStdMsg(srcFile string, msg string) {
	panic(fmt.Sprintf("[Error (%v)]: %v\n", srcFile, msg))
}
