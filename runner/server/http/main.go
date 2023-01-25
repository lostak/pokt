package main

import (
	"fmt"

	"github.com/lostak/pokt/server"
)

const port = ":3333"

func main() {
	err := server.ServeHTTP(port)
	if err != nil {
		fmt.Println("Server Error: ", err.Error())
	}
}
