// +build windows

package userdirs

import (
	"path/filepath"

	"github.com/apparentlymart/go-userdirs/windowsbase"
)

func supportedOS() bool {
	return true
}

func forApp(name string, vendor string, bundleID string) Dirs {
	subDir := filepath.Join(vendor, name)
	roamingDir := filepath.Join(windowsbase.RoamingAppDataDir(), subDir)
	localDir := filepath.Join(windowsbase.LocalAppDataDir(), subDir)

	return Dirs{
		ConfigDirs: []string{roamingDir},
		DataDirs:   []string{roamingDir},
		CacheDir:   localDir,
	}
}
