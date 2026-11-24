// Package version holds the build identity shared by the three Nextar binaries.
package version

import "fmt"

// Version, Commit and Date are set at build time with -ldflags -X. The defaults
// identify a binary built with plain `go build`, which must never be mistaken
// for a release.
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

// String returns the one-line identity printed by `<binary> version`.
func String(binary string) string {
	return fmt.Sprintf("%s %s (commit %s, built %s)", binary, Version, Commit, Date)
}
