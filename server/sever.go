package server

import (
	"fmt"
	"net/http"
)

func ServeHTTP(port string) error {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/", fs)

	fmt.Printf("Listening on http://localhost:%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}
}
