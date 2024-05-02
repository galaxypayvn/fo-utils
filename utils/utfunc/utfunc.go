package utfunc

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"log"
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

// Func to get random string using snowflake algorithm
func GetRandomString() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}

	// Generate a snowflake ID.
	id := node.Generate()

	return id.String()
}
