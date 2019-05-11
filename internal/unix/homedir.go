package unix

import (
	"os"
	"os/user"
	"path/filepath"
)

// Home returns the home directory for the current process, with the following
// preference order:
//
//     - The value of the HOME environment variable, if it is set and contains
//       an absolute path.
//     - The home directory indicated in the return value of the "Current"
//       function in the os/user standard library package, which has
//       platform-specific behavior, if it contains an absolute path.
//     - If neither of the above yields an absolute path, the string "/".
//
// In practice, POSIX requires the HOME environment variable to be set, so on
// any reasonable system it is that which will be selected. The other
// permutations are fallback behavior for less reasonable systems.
//
// XDG does not permit applications to write directly into the home directory.
// Instead, the paths returned by other functions in this package are
// potentially derived from the home path, if their explicit environment
// variables are not set.
func Home() string {
	if homeDir := os.Getenv("HOME"); homeDir != "" {
		if filepath.IsAbs(homeDir) {
			return homeDir
		}
	}

	user, err := user.Current()
	if err != nil {
		if homeDir := user.HomeDir; homeDir != "" {
			if filepath.IsAbs(homeDir) {
				return homeDir
			}
		}
	}

	// Fallback behavior mimics common choice in other software.
	return "/"
}
