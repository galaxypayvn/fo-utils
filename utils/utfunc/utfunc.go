package utfunc

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"bitbucket.org/finesys/finesys-utility/libs/serror"
)

// Try function
func Try(fn func() serror.SError) (errx serror.SError) {
	if fn == nil {
		return errx
	}

	func() {
		defer func() {
			if errx != nil {
				return
			}

			if errRcv := recover(); errRcv != nil {
				errx = serror.Newsc(1, fmt.Sprintf("%+v", errRcv), "Unexpected exception has occurred")
			}
		}()

		errx = fn()
	}()

	return errx
}

func CurrentFunctionName(level int) string {
	pc, _, _, _ := runtime.Caller(1 + level)
	strArr := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	if len(strArr) > 0 {
		return strArr[len(strArr)-1]
	}
	return ""
}

func GetCurrentCaller(caller interface{}, level int) string {
	strArr := strings.Split(reflect.TypeOf(caller).String(), ".")
	if caller != nil && len(strArr) > 0 {
		return fmt.Sprintf("%s.%s", strArr[len(strArr)-1], CurrentFunctionName(1+level))
	}
	return CurrentFunctionName(1)
}
