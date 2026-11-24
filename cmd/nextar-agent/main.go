// Command nextar-agent is the agent that runs on every registered worker.
package main

import (
	"os"

	"github.com/gsoultan/nextar-server/internal/version"
)

func main() {
	os.Exit(version.Run("nextar-agent", os.Args[1:], os.Stdout, os.Stderr))
}
