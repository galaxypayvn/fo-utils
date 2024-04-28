package utfunc

import (
	"fmt"
	"runtime"
	"strings"
)

// GetCurrentCaller returns the name of the current caller function.
func GetCurrentCaller(caller interface{}, level int) string {
	if caller != nil {
		callerType := fmt.Sprintf("%T", caller)
		parts := strings.Split(callerType, ".")
		if len(parts) > 0 {
			return fmt.Sprintf("%s.%s", parts[len(parts)-1], getCurrentFunctionName(level+1))
		}
	}
	return getCurrentFunctionName(level + 1)
}

func getCurrentFunctionName(level int) string {
	pc, _, _, _ := runtime.Caller(level + 1)
	name := runtime.FuncForPC(pc).Name()
	parts := strings.Split(name, ".")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}
