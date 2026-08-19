package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(w, "Task api is running ")

	})
	fmt.Println(" server is runnuing on http://localhost:8000")

	http.ListenAndServe(":8000", nil)
}
