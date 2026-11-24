package version

import (
	"fmt"
	"io"
)

// Run is the whole behaviour of a binary until its real command is written:
// it answers `version`, and refuses everything else with a non-zero exit code
// so that a service manager never reports a placeholder as running.
func Run(binary string, args []string, stdout, stderr io.Writer) int {
	if wantsVersion(args) {
		fmt.Fprintln(stdout, String(binary))
		return 0
	}
	fmt.Fprintf(stderr, "%s: not implemented yet; only `%s version` works in this build\n", binary, binary)
	return 1
}

func wantsVersion(args []string) bool {
	if len(args) != 1 {
		return false
	}
	return args[0] == "version" || args[0] == "--version"
}
