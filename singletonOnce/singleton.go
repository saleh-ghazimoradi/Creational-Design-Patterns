package main

import "fmt"

type single struct{}

var singleton *single

func init() {
	fmt.Println("There must be one and only one instance of an object!!!")
	singleton = &single{}
}

func getInstance() *single {
	return singleton
}
