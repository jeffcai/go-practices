package main

import (
	"os"

	"github.com/jeffcai/go-practice/pkg/hello"
	"github.com/spf13/pflag"
)

func main() {
	flags := pflag.NewFlagSet("hello-golang", pflag.ExitOnError)
	pflag.CommandLine = flags

	// bypass to DebugCmd
	cmd := hello.GolangCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
