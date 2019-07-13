package main

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/james-woo/wakeus/api/controllers"
	"github.com/james-woo/wakeus/api/jobs"
	"github.com/james-woo/wakeus/api/rpc"
	"log"
	"net/http"
	"os"
)

// env GOOS=linux GOARCH=arm GOARM=5 go build
func main() {
	router := mux.NewRouter()
	// Tasks
	router.HandleFunc("/api/task", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{task_id:[0-9]+}", controllers.GetTask).Methods("GET")
	router.HandleFunc("/api/tasks/{task_id:[0-9]+}", controllers.UpdateTask).Methods("PATCH")
	router.HandleFunc("/api/tasks/{task_id:[0-9]+}", controllers.DeleteTask).Methods("DELETE")

	// Perform
	router.HandleFunc("/api/command/basic", controllers.Basic).Methods("POST")
	router.HandleFunc("/api/command/fade", controllers.Fade).Methods("POST")
	router.HandleFunc("/api/command/rainbow", controllers.Rainbow).Methods("POST")
	router.HandleFunc("/api/command/clear", controllers.Clear).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Launch periodic tasks
	fmt.Printf("Launching tasks\n")
	jobs.LaunchJobs(context.Background())

	//Perform hardware test
	fmt.Printf("Performing hardware test\n")
	_, _ = rpc.PerformTest(context.Background())

	// Launch app, visit localhost:8000/api
	fmt.Printf("Running on localhost:%s\n", port)
	handler := handlers.CORS(
		handlers.AllowedHeaders([]string{"Accept", "Authorization", "Content-Type", "Content-Length", "Accept-Encoding", "X-Requested-With"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)
	log.Fatal(http.ListenAndServe(":" + port, handler(router)))
}
