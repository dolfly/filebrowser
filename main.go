package main

import (
	"runtime"

	"github.com/dolfly/filebrowser/v2/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
