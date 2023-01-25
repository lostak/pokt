package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Go web assembly")

	// Get js.Funcs to be supported
	js.Global().Set("formatPortfolioJSON", GetPortfolioWrapper())
	fmt.Println("formatPortfolioJSON set to be the js.Func returned from GetPortfolioWrapper()")

	js.Global().Set("formatAccountJSON", GetAccountWrapper())
	fmt.Println("formatAccountJSON set to be the js.Func returned from GetAccountWrapper()")

	<-make(chan bool)
}
