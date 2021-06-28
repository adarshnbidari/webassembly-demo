package main

import (
	"fmt"
	"syscall/js"
)

func main() {

	c := make(chan bool)

	fmt.Println("This is a main function")

	js.Global().Set("adarsh", js.FuncOf(test))

	c <- true

}

func test(this js.Value, inputs []js.Value) interface{} {

	fmt.Println("Testing exported function 😁")

	fmt.Println("Hello", inputs[0].String())

	return nil

}
