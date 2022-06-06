package functioncaller

import (
	"fmt"
	"runtime"
)

func PrintFuncName() string {
	fpcs := make([]uintptr, 1)

	// Skip 3 levels to get to the caller of whoever called Caller()
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		fmt.Println("MSG: Couldn't get caller information")
	}

	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		fmt.Println("MSG: Couldn't get caller function")
	}

	if fun == nil {
		return "unknown"
	}

	return fun.Name()
}
