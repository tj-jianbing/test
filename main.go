package main

import "fmt"

type conncetion interface {
	Name() string
	Open()
	Exec()
}

type Phone struct {
	name string
	Note string
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
	v1 = Phone{"T1", "good"}
	v2 := Phone{"T1", "good"}
	v1.Open()

	fmt.Println(v2.name)
	fmt.Println(v2.Note)

	fmt.Println(v1.Name())
	fmt.Println("hello world!")
}
