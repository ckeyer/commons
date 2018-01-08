package format

import "fmt"

const (
	cRedBegin   = "\033[31m"
	cGreenBegin = "\033[32m"
	cBrownBegin = "\033[33m"
	cEnd        = "\033[0m"
)

func ErrorString(format string, args ...interface{}) string {
	return cRedBegin + fmt.Sprintf(format, args...) + cEnd
}

func SuccessString(format string, args ...interface{}) string {
	return cGreenBegin + fmt.Sprintf(format, args...) + cEnd
}

func WarnString(format string, args ...interface{}) string {
	return cBrownBegin + fmt.Sprintf(format, args...) + cEnd
}
