package userdirs

import (
	"os"
)

func testTempEnv(name, value string) func() {
	old := os.Getenv(name)
	os.Setenv(name, value)
	return func() {
		os.Setenv(name, old)
	}
}

func testTempEnvMany(vars map[string]string) func() {
	old := make(map[string]string, len(vars))
	for name, value := range vars {
		old[name] = os.Getenv(name)
		os.Setenv(name, value)
	}
	return func() {
		for name, value := range old {
			os.Setenv(name, value)
		}
	}
}
