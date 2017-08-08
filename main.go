package main

import "fmt"

type conncetion interface {
	Name() string
	Open()
	Exec()
}

type Phone struct {
	name string
}

func (con Phone) Name() string {
	return con.name
}

func (con Phone) Open() {
	fmt.Println("test open interface")
}

func (con Phone) Exec() {}

func main() {
	var v1 conncetion
	v1 = Phone{"T1"}
	v1.Open()

	fmt.Println("hello world!")
}
