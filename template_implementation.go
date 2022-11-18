package main

import "fmt"

type hello struct {
	world string
}

func (h hello) hello() {
	fmt.Println("hello " + h.world)
}