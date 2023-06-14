package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleton *single

func getInstance() *single {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			fmt.Println("1. There must be one and only one instance of an object!!!")
			singleton = &single{}
		} else {
			fmt.Println("2. The instance had already been created!!!")
		}
	} else {
		fmt.Println("3. This one as well")
	}
	return singleton
}
