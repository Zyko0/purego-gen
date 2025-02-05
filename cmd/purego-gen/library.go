package main

type Library struct {
	PathByOS map[string]string
	Alias    string
}

type Symbol struct {
	Name   string
	Symbol string
}
