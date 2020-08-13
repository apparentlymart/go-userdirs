// +build linux aix dragonfly freebsd netbsd openbsd solaris

package userdirs

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestForAppUnix(t *testing.T) {
	tests := map[string]struct {
		env  map[string]string
		want Dirs
	}{
		"defaults": {
			map[string]string{
				"HOME":            "/home/placeholder",
				"XDG_DATA_HOME":   "",
				"XDG_DATA_DIRS":   "",
				"XDG_CONFIG_HOME": "",
				"XDG_CONFIG_DIRS": "",
				"XDG_CACHE_HOME":  "",
			},
			Dirs{
				ConfigDirs: []string{
					"/home/placeholder/.config/spline-reticulator",
					"/etc/xdg/spline-reticulator",
				},
				DataDirs: []string{
					"/home/placeholder/.local/share/spline-reticulator",
					"/usr/local/share/spline-reticulator",
					"/usr/share/spline-reticulator",
				},
				CacheDir: "/home/placeholder/.cache/spline-reticulator",
			},
		},
		"overridden": {
			map[string]string{
				"HOME":            "/home/placeholder",
				"XDG_DATA_HOME":   "/somewhere/else",
				"XDG_DATA_DIRS":   "/another/place:/yet/another/place",
				"XDG_CONFIG_HOME": "/primary/config",
				"XDG_CONFIG_DIRS": "/more/config:/even/more/config",
				"XDG_CACHE_HOME":  "/cache/over/here",
			},
			Dirs{
				ConfigDirs: []string{
					"/primary/config/spline-reticulator",
					"/more/config/spline-reticulator",
					"/even/more/config/spline-reticulator",
				},
				DataDirs: []string{
					"/somewhere/else/spline-reticulator",
					"/another/place/spline-reticulator",
					"/yet/another/place/spline-reticulator",
				},
				CacheDir: "/cache/over/here/spline-reticulator",
			},
		},
		"invalid relative paths": {
			map[string]string{
				"HOME": "/home/placeholder",
				// Relative paths are not permitted and are ignored
				"XDG_DATA_HOME":   "bar",
				"XDG_DATA_DIRS":   "/valid-data:baz",
				"XDG_CONFIG_HOME": "boop",
				"XDG_CONFIG_DIRS": "beep:/valid-config",
				"XDG_CACHE_HOME":  "blarp",
			},
			Dirs{
				ConfigDirs: []string{
					"/home/placeholder/.config/spline-reticulator",
					"/valid-config/spline-reticulator",
				},
				DataDirs: []string{
					"/home/placeholder/.local/share/spline-reticulator",
					"/valid-data/spline-reticulator",
				},
				CacheDir: "/home/placeholder/.cache/spline-reticulator",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer testTempEnvMany(test.env)()
			got := ForApp("Spline Reticulator", "Acme Corp", "com.example.splines")

			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("wrong result\n%s", diff)
			}
		})
	}
}

func TestForAppUnixNoEnv(t *testing.T) {
	// This test deals with the situation where no environment variables are
	// set at all, including $HOME. In that case the results will vary depending
	// on where we are running, so we'll run this one only inside the Travis-CI
	// test runs so we can depend on its execution environment.
	if travisEnv := os.Getenv("TRAVIS"); travisEnv == "" {
		t.Skipf("No-environment tests run only in Travis-CI")
	}

	defer testTempEnvMany(map[string]string{
		"HOME":            "",
		"XDG_DATA_HOME":   "",
		"XDG_DATA_DIRS":   "",
		"XDG_CONFIG_HOME": "",
		"XDG_CONFIG_DIRS": "",
		"XDG_CACHE_HOME":  "",
	})()

	got := ForApp("Spline Reticulator", "Acme Corp", "com.example.splines")
	want := Dirs{
		ConfigDirs: []string{
			"/home/travis/.config/spline-reticulator",
			"/etc/xdg/spline-reticulator",
		},
		DataDirs: []string{
			"/home/travis/.local/share/spline-reticulator",
			"/usr/local/share/spline-reticulator",
			"/usr/share/spline-reticulator",
		},
		CacheDir: "/home/travis/.cache/spline-reticulator",
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("wrong result\n%s", diff)
	}
}
