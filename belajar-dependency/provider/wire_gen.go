// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

// Injectors from injector.go:

func InitializedService() *SimpleService {
	simpleRepository := NewSimpleRepository()
	simpleService := NewSimpleService(simpleRepository)
	return simpleService
}
