package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func M() {
	fmt.Println("Dookie")
}

func main() {
	var i I = T{"hello"}
	i.M()
	M()
}
