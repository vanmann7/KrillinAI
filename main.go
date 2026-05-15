package main

import (
	"fmt"
	"os"

	"github.com/krillinai/KrillinAI/cmd"
)

// Version information injected at build time via ldflags
var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

func main() {
	// Pass build-time version info into the command layer
	cmd.SetVersionInfo(Version, Commit, BuildDate)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		// Use exit code 2 to distinguish application errors from OS-level errors (code 1)
		os.Exit(2)
	}
}
