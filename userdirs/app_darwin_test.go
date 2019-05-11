// +build darwin

package userdirs

import (
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
