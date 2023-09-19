// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/dragon2org/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	context2 "context"
)

// Injectors from wireg.go:

func inject(contextContext context2.Context, arg struct{}) (context, error) {
	mainContext, err := provide(contextContext)
	if err != nil {
		return context{}, err
	}
	return mainContext, nil
}
