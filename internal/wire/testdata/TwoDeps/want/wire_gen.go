// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/dragon2org/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFooBar() FooBar {
	foo := provideFoo()
	bar := provideBar()
	fooBar := provideFooBar(foo, bar)
	return fooBar
}
