package main

import (
	"fmt"
	"reflect"
	"time"
)

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) incrementAge(inc int) {
	a.Age += inc
}

func main() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		for range [10]int{} {
			one <- "one"
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for range [10]int{} {
			two <- "two"
			time.Sleep(time.Second * 2)
		}
	}()

	for range [20]int{} {
		select {
		case num1 := <-one:
			fmt.Println(num1)
		case num2 := <-two:
			fmt.Println(num2)
		}
	}
}

func doSomething(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(v)
}
