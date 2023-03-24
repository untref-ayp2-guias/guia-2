package main

import (
	"fmt"
	"guia2/stack"
)

func main() {
	var s stack.Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	var s2 stack.Stack
	s2.Push("Hola")
	s2.Push("Mundo")
	fmt.Println(s2.Pop())
	fmt.Println(s2.Pop())
	fmt.Println(s2.Pop())

	var s3 stack.Stack
	s3.Push(1)
	s3.Push(2)

	//Modificamos la pila sin usar los métodos definidos
	s3[0] = "Hola"
	s3[1] = "Mundo"
	fmt.Println(s3.Pop())
	fmt.Println(s3.Pop())
	fmt.Println(s3.Pop())

}
