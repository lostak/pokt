package main

import (
	"fmt"

	"github.com/lostak/pokt/server"
)

func main() {
	err := server.ServeHTTP(":3333")
	if err != nil {
		fmt.Println("Server Error: ", err.Error())
	}
}
