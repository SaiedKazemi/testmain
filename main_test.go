package main

import (
	"flag"
	"os"
	"testing"
)

var (
	forcedFailure = flag.Bool("forcefailure", false, "Run the tests that force failures.")

	tests = []struct {
		name       string
		skipTest   bool
		cmdArgs    []string
		wantOutput string
	}{
		{"valid argument", false, []string{"--arg", "valid"}, ""},
		{"forced failure 1", true, []string{"--arg", "valid"}, "forced failure"},
		{"invalid argument", false, []string{"--arg", "invalid"}, "invalid argument: invalid"},
		{"forced failure 2", true, []string{"--arg", "invalid"}, ""},
	}
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestMainFunc(t *testing.T) {
	for _, test := range tests {
		test := test // not running parallel tests, but good to do this anyway
		if test.skipTest && !*forcedFailure {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			if gotOutput := runMain(test.cmdArgs); gotOutput != test.wantOutput {
				t.Errorf("main(%v) = %q, want %q", test.cmdArgs, gotOutput, test.wantOutput)
			}
		})
	}
}

func runMain(cmdArgs []string) string {
	oldArgs := os.Args
	oldPanicOnError := panicOnError
	defer func() {
		os.Args = oldArgs
		panicOnError = oldPanicOnError
	}()
	os.Args = append([]string{"main"}, cmdArgs...)
	panicOnError = true
	main()
	return fatalError
}
