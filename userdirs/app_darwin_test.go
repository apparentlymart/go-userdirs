// +build darwin

package userdirs

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestForAppDarwin(t *testing.T) {
	tests := map[string]struct {
		env  map[string]string
		want Dirs
	}{
		"defaults": {
			map[string]string{
				"HOME":            "/Users/placeholder",
			},
			Dirs{
				ConfigDirs: []string{
					"/Users/placeholder/Library/Application Support/com.example.splines",
				},
				DataDirs: []string{
					"/Users/placeholder/Library/Application Support/com.example.splines",
					"/Library/Application Support/com.example.splines",
				},
				CacheDir: "/Users/placeholder/Library/Caches/com.example.splines",
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

func TestForAppDarwinNoEnv(t *testing.T) {
	// This test deals with the situation where no environment variables are
	// set at all, including $HOME. In that case the results will vary depending
	// on where we are running, so we'll run this one only inside the Travis-CI
	// test runs so we can depend on its execution environment.
	if travisEnv := os.Getenv("TRAVIS"); travisEnv == "" {
		t.Skipf("No-environment tests run only in Travis-CI")
	}

	defer testTempEnvMany(map[string]string{
		"HOME":            "",
	})()

	got := ForApp("Spline Reticulator", "Acme Corp", "com.example.splines")
	want := Dirs{
		ConfigDirs: []string{
			"/Users/travis/Library/Application Support/com.example.splines",
		},
		DataDirs: []string{
			"/Users/travis/Library/Application Support/com.example.splines",
			"/Library/Application Support/com.example.splines",
		},
		CacheDir: "/Users/travis/Library/Caches/com.example.splines",
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("wrong result\n%s", diff)
	}
}
