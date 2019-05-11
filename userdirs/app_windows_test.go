// +build windows

package userdirs

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestForAppWindows(t *testing.T) {
	// There are no environment variable mechanisms for us to override this
	// on Windows (the Windows Registry is the source of record) so for
	// simplicity right now we're only running these tests in the Travis-CI
	// environment, where we can depend on a particular configuration.
	//
	// (If you're developing on a Windows system and want to run this anyway,
	// you can set the TRAVIS environment variable to force it. The paths
	// probably won't match what is in this test, but you can review them
	// manually and see that they follow the similar naming pattern, at least.)
	if travisEnv := os.Getenv("TRAVIS"); travisEnv == "" {
		t.Skipf("Windows tests run only in Travis-CI")
	}

	got := ForApp("Spline Reticulator", "Acme Corp", "com.example.splines")
	want := Dirs{
		ConfigDirs: []string{`C:\Users\travis\AppData\Roaming\Acme Corp\Spline Reticulator`},
		DataDirs:   []string{`C:\Users\travis\AppData\Roaming\Acme Corp\Spline Reticulator`},
		CacheDir:   `C:\Users\travis\AppData\Local\Acme Corp\Spline Reticulator`,
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("wrong result\n%s", diff)
	}
}
