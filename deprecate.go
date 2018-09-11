package oncer

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

const deprecated = "DEPRECATED"

var deprecationWriter io.Writer = os.Stdout

func Deprecate(depth int, name string, msg string) {
	Do(deprecated+name, func() {
		if depth <= 0 {
			depth = 5
		}
		// for i := 0; i < 10; i++ {
		// 	_, file, line, _ := runtime.Caller(i)
		// 	fmt.Println("### i ->", i)
		// 	fmt.Println("### file ->", file)
		// 	fmt.Println("### line ->", line)
		// }
		_, file, line, _ := runtime.Caller(depth)
		fmt.Fprintf(deprecationWriter, "[%s] %s has been deprecated. (%s:%d)\n", deprecated, name, file, line)
		if len(msg) > 0 {
			fmt.Fprintf(deprecationWriter, "\t%s\n", msg)
		}
	})
}
