package userdirs

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDirsNewPath(t *testing.T) {
	dirs := Dirs{
		ConfigDirs: []string{
			"/user/config",
			"/global/config",
		},
		DataDirs: []string{
			"/user/data",
			"/global/data",
		},
		CacheDir: "/cache",
	}

	t.Run("NewConfigPath", func(t *testing.T) {
		got := dirs.NewConfigPath("foo", "bar")
		want := filepath.Join("/user", "config", "foo", "bar")
		if got != want {
			t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
		}
	})
	t.Run("NewDataPath", func(t *testing.T) {
		got := dirs.NewDataPath("foo", "bar")
		want := filepath.Join("/user", "data", "foo", "bar")
		if got != want {
			t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
		}
	})
	t.Run("CachePath", func(t *testing.T) {
		got := dirs.CachePath("foo", "bar")
		want := filepath.Join("/cache", "foo", "bar")
		if got != want {
			t.Errorf("wrong result\ngot:  %s\nwant: %s", got, want)
		}
	})
}

func TestDirsFindConfigFiles(t *testing.T) {
	t.Run("both", func(t *testing.T) {
		prefix, dirs := testdataDirs(t)
		got := dirs.FindConfigFiles("foo.conf")
		want := []string{
			filepath.Join(prefix, "user", "config", "foo.conf"),
			filepath.Join(prefix, "global", "config", "foo.conf"),
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("wrong result\n%s", diff)
		}
	})
	t.Run("only one", func(t *testing.T) {
		prefix, dirs := testdataDirs(t)
		got := dirs.FindConfigFiles("bar.conf")
		want := []string{
			filepath.Join(prefix, "global", "config", "bar.conf"),
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("wrong result\n%s", diff)
		}
	})
	t.Run("none", func(t *testing.T) {
		_, dirs := testdataDirs(t)
		got := dirs.FindConfigFiles("nonexist.conf")
		want := []string(nil)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("wrong result\n%s", diff)
		}
	})
}

func TestDirsFindDataFiles(t *testing.T) {
	t.Run("both", func(t *testing.T) {
		prefix, dirs := testdataDirs(t)
		got := dirs.FindDataFiles("baz.dat")
		want := []string{
			filepath.Join(prefix, "user", "data", "baz.dat"),
			filepath.Join(prefix, "global", "data", "baz.dat"),
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("wrong result\n%s", diff)
		}
	})
	t.Run("none", func(t *testing.T) {
		_, dirs := testdataDirs(t)
		got := dirs.FindDataFiles("nonexist.dat")
		want := []string(nil)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("wrong result\n%s", diff)
		}
	})
}

func TestDirsGlobConfigFiles(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		prefix, dirs := testdataDirs(t)
		got := dirs.GlobConfigFiles("*.conf")
		want := []string{
			filepath.Join(prefix, "user", "config", "foo.conf"),
			filepath.Join(prefix, "global", "config", "bar.conf"),
			filepath.Join(prefix, "global", "config", "foo.conf"),
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("wrong result\n%s", diff)
		}
	})
	t.Run("none", func(t *testing.T) {
		_, dirs := testdataDirs(t)
		got := dirs.GlobConfigFiles("*.nope")
		want := []string(nil)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("wrong result\n%s", diff)
		}
	})
}

func TestDirsGlobDataFiles(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		prefix, dirs := testdataDirs(t)
		got := dirs.GlobDataFiles("*.dat")
		want := []string{
			filepath.Join(prefix, "user", "data", "baz.dat"),
			filepath.Join(prefix, "global", "data", "baz.dat"),
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("wrong result\n%s", diff)
		}
	})
	t.Run("none", func(t *testing.T) {
		_, dirs := testdataDirs(t)
		got := dirs.GlobDataFiles("*.nope")
		want := []string(nil)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("wrong result\n%s", diff)
		}
	})
}

func testdataDirs(t *testing.T) (prefix string, dirs Dirs) {
	prefix, err := filepath.Abs("testdata")
	if err != nil {
		t.Fatalf("cannot find absolute path for testdata dir: %s", err)
	}

	return prefix, Dirs{
		ConfigDirs: []string{
			filepath.Join(prefix, "user", "config"),
			filepath.Join(prefix, "global", "config"),
		},
		DataDirs: []string{
			filepath.Join(prefix, "user", "data"),
			filepath.Join(prefix, "global", "data"),
		},
		CacheDir: filepath.Join(prefix, "cache"),
	}
}
