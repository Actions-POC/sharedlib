// Package sharedlib a package for developers who prefer github
package sharedlib

const currentVersion = "1.0.0"

// Version returns the library name, please change it for every release. Just
// for to show changes over time
func Version() string {
	return currentVersion
}

// Name returns the library name
func Name() string {
	return "sharedlib"
}

// Hi let's be friends
func Hi(name string) string {
	return "👋  Hi " + name
}