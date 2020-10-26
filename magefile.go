// +build mage

package main

import "github.com/magefile/mage/sh"

// Build generates a binary of the project
func Build() error {
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}

	return sh.Run("go", "build", "-o", "bin/scarecrow", "--ldflags", "-s -w", "./")
}

// Format lints and fixes all files in the directory
func Format() error {
	return sh.Run("go", "fmt", "./...")
}
