package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/james-woo/wakeus/api/app"
	"github.com/james-woo/wakeus/api/controllers"
	"github.com/james-woo/wakeus/api/jobs"
	"github.com/james-woo/wakeus/api/rpc"
	"net/http"
	"os"
)

// env GOOS=linux GOARCH=arm GOARM=5 go build
func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // Attach JWT auth middleware

	// Tasks
	router.HandleFunc("/api/task", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{task_id:[0-9]+}", controllers.GetTask).Methods("GET")
	router.HandleFunc("/api/tasks/{task_id:[0-9]+}", controllers.UpdateTask).Methods("PATCH")
	router.HandleFunc("/api/tasks/{task_id:[0-9]+}", controllers.DeleteTask).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Launch periodic tasks
	fmt.Printf("Launching tasks\n")
	jobs.LaunchTasks(context.Background())

	//Perform hardware test
	fmt.Printf("Performing hardware test\n")
	rpc.PerformTest(context.Background())

	fmt.Printf("Performing hardware basic\n")
	rpc.PerformBasic(
		context.Background(),
		rpc.Color{R: 25, G: 25, B: 25},
		1,
	)

	fmt.Printf("Performing hardware fade\n")
	rpc.PerformFade(
		context.Background(),
		rpc.Color{R: 0, G: 0, B: 0},
		rpc.Color{R: 255, G: 255, B: 255},
		0,
		1,
		5000,
	)

	fmt.Printf("Performing hardware clear\n")
	rpc.PerformClear(context.Background())

	// Launch app, visit localhost:8000/api
	fmt.Printf("Running on localhost:%s\n", port)
	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}
