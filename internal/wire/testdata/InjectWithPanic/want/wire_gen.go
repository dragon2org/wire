// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/dragon2org/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wireg.go:

func injectedMessage() string {
	string2 := provideMessage()
	return string2
}
