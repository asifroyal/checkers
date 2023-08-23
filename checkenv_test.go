package main

import (
	"os"
	"testing"
)

func TestCheckEnv(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		directories    []string
		extensions     []string
		ignoreDirs     []string
		expectedOutput string
		setup          func()
		teardown       func()
	}{
		{
			name:           "missing environment variable",
			directories:    []string{"testdata"},
			extensions:     []string{".js"},
			ignoreDirs:     []string{},
			expectedOutput: "Missing variable: TEST_VAR\n",
		},
		{
			name:           "environment variable set",
			directories:    []string{"testdata"},
			extensions:     []string{".js"},
			ignoreDirs:     []string{},
			expectedOutput: "",
			setup: func() {
				os.Setenv("TEST_VAR", "test")
			},
			teardown: func() {
				os.Unsetenv("TEST_VAR")
			},
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up test case
			if tc.setup != nil {
				tc.setup()
			}
			defer func() {
				if tc.teardown != nil {
					tc.teardown()
				}
			}()

			// Run function and capture output
			output := captureOutput(func() {
				checkEnv(tc.directories, tc.extensions, tc.ignoreDirs, scanFiles)
			})

			// Check output
			if output != tc.expectedOutput {
				t.Errorf("Expected output: %s, but got: %s", tc.expectedOutput, output)
			}
		})
	}
}

// captureOutput captures the output of a function
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	output := make(chan string)
	go func() {
		buf := make([]byte, 1024)
		n, _ := r.Read(buf)
		output <- string(buf[:n])
	}()

	return <-output
}
