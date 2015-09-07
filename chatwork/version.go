package chatwork

import (
	"fmt"
	"runtime"
)

const VERSION = "0.0.1"

func PrintVersion() {
	fmt.Printf(`Chatwork agent %s
Compiler: %s %s
Copyright (C) 2015 Shinichiro Yuki <sinycourage@gmail.com>
`,
		VERSION,
		runtime.Compiler,
		runtime.Version())
}
