package main

import (
	"fmt"
	"net/http"
	"encoding/json"
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

	// http.HandleFunc("/tasks", getTask)
	// http.HandlerFunc("/tasks", createTask)
	http.HandleFunc("/tasks", tasksHandler)
	fmt.Println(" server is runnuing on http://localhost:8000")

	http.ListenAndServe(":8000", nil)

}
func getTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}


func createTask( w http.ResponseWriter, r *http.Request) {

	var newtask Task 

	err := json.NewDecoder(r.Body).Decode(&newtask)

	if err != nil{

		http.Error(w ,"invalid request body", http.StatusBadRequest)
		return

	}

	newtask.ID = len(tasks) + 1

	tasks = append(tasks, newtask)

	w.Header().Set("Content-type", "application/json")

json.NewEncoder(w).Encode(newtask)


}

func tasksHandler(w http.ResponseWriter, r *http.Request) {


	if r.Method ==http.MethodGet{
getTask(w,r )
return


	}

	if r.Method == http.MethodPost{

		createTask(w,r)
		return
	}


	http.Error(w, "methods are not allowed", http.StatusMethodNotAllowed)
}