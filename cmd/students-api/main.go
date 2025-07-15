package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ashunasar/golang-students-crud-api/internal/config"
	"github.com/ashunasar/golang-students-crud-api/internal/http/handlers/student"
	"github.com/ashunasar/golang-students-crud-api/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}
	slog.Info("storage Initialized", slog.String("Env", cfg.Env))

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students", student.GetStudents(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetStudentById(storage))
	router.HandleFunc("PUT /api/students", student.UpdateStudent(storage))
	router.HandleFunc("DELETE /api/students/{id}", student.DeleteStudent(storage))

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	slog.Info("Server started", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server Shutdown successfully ")

}
