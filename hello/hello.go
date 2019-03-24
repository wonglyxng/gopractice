package main

import "fmt"

type A struct {
	value int
	name string
}

func (a *A)f()  {
	a.name = "a name"
	a.value = 1
}

func main() {
	var a A
	a.f()
	fmt.Printf("type=%T, a=%#v",a,a)
}
