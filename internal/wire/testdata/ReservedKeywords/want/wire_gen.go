// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/dragon2org/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectInterface() Interface {
	select2 := provideSelect()
	mainInterface := provideInterface(select2)
	return mainInterface
}

// wire.go:

// Wire tries to disambiguate the variable "select" by prepending
// the package name; this package-scoped variable conflicts with that
// and forces a different name.
var mainSelect = 0
