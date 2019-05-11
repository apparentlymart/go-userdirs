package userdirs

// Dirs represents a set of directory paths with different purposes.
type Dirs struct {
	// ConfigDirs is a list, in preference order, of directory paths to search
	// for configuration files.
	//
	// The list must always contain at least one element, and its first element
	// is the directory where any new configuration files should be written.
	//
	// On some systems, ConfigDirs and DataDirs may overlap, so applications
	// which scan the contents of the configuration directories should impose
	// some additional filtering to distinguish configuration files from data
	// files.
	ConfigDirs []string

	// DataDirs is a list, in preference order, of directory paths to search for
	// data files.
	//
	// The list must always contain at least one element, and its first element
	// is the directory where any new data files should be written.
	//
	// On some systems, ConfigDirs and DataDirs may overlap, so applications
	// which scan the contents of the data directories should impose some
	// additional filtering to distinguish data files from configuration files.
	DataDirs []string

	// CacheDir is the path of a single directory that can be used for temporary
	// cache data.
	//
	// The cache is suitable only for data that the calling application could
	// recreate if lost. Any file or directory under this prefix may be deleted
	// at any time by other software.
	//
	// This directory may, on some systems, match one of the directories
	// returned in ConfigDirs and/or DataDirs. For this reason applications
	// must ensure that they do not misinterpret config and data files as
	// cache files, and in particular should not naively purge a cache by
	// emptying this directory.
	CacheDir string
}

// ConfigHome returns the path for the directory where any new configuration
// files should be written.
func (d Dirs) ConfigHome() string {
	return d.ConfigDirs[0]
}

// DataHome returns the path for the directory where any new configuration
// files should be written.
func (d Dirs) DataHome() string {
	return d.DataDirs[0]
}
