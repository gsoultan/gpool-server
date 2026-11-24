// Command nextar-server is the manager: UI, API and state store.
package main

import (
	"os"

	"github.com/gsoultan/nextar-server/internal/version"
)

func main() {
	os.Exit(version.Run("nextar-server", os.Args[1:], os.Stdout, os.Stderr))
}
