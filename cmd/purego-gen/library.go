package main

import "github.com/Zyko0/purego-gen/internal"

type Library struct {
	PathByOS *internal.OrderedMap[string, string]
	Alias    string
}

type Symbol struct {
	Name   string
	Symbol string
}
