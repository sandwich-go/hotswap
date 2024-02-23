package main

import (
	"os"

	"github.com/sandwich-go/hotswap/cli/hotswap/cmd"
	"github.com/sandwich-go/hotswap/cli/hotswap/g"
)

func main() {
	for i, arg := range os.Args {
		if arg == "--" {
			g.BuildFlags = os.Args[i+1:]
			os.Args = os.Args[:i]
			break
		}
	}

	cmd.Execute()
}
