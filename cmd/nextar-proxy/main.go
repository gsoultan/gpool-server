// Command nextar-proxy is the proxy in front of the database on a worker.
package main

import (
	"os"

	"github.com/gsoultan/nextar-server/internal/version"
)

func main() {
	os.Exit(version.Run("nextar-proxy", os.Args[1:], os.Stdout, os.Stderr))
}
