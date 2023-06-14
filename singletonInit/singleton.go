package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct{}

var singleton *single

func getInstance() *single {
	if singleton == nil {
		once.Do(func() {
			fmt.Println("There must be one and only one instance of an object!!!")
			singleton = &single{}
		})
	} else {
		fmt.Println("The other one already... ")
	}
	return singleton
}
