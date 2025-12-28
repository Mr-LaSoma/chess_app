package errors

import "fmt"

// PanicStdMsg is a method created to panic the program with a standard format.
// | Error | srcFile name (with the parent folder) | error message
func PanicStdMsg(srcFile string, msg string) {
	panic(fmt.Sprintf("[Error (%v)]: %v\n", srcFile, msg))
}
