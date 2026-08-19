package main

import (
	"fmt"
	"net/http"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks = []Task{
	{

		ID:        1,
		Title:     "learn go",
		Completed: false,
	},
	{
		ID:        2,
		Title:     "Build API",
		Completed: false,
	},
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "Task api is running ")

	})

	http.HandleFunc("/tasks", getTask)
	fmt.Println(" server is runnuing on http://localhost:8000")

	http.ListenAndServe(":8000", nil)

}
func getTask(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, tasks)
}
