//go:build darwin || freebsd || linux

package puregogen

import "github.com/ebitengine/purego"

func OpenLibrary(name string) (uintptr, error) {
	return purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}

func OpenSymbol(lib uintptr, name string) (uintptr, error) {
	return purego.Dlsym(lib, name)
}