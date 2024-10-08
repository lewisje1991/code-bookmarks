package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"log/slog"

	appdiary "github.com/lewisje1991/code-bookmarks/internal/app/diary"
	domaindiary "github.com/lewisje1991/code-bookmarks/internal/domain/diary"

	apptasks "github.com/lewisje1991/code-bookmarks/internal/app/tasks"
	domaintasks "github.com/lewisje1991/code-bookmarks/internal/domain/tasks"

	"github.com/lewisje1991/code-bookmarks/internal/foundation/config"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/postgres"
	"github.com/lewisje1991/code-bookmarks/internal/foundation/server"
)

// TODO: authorization/RBAC
// TODO: use supabase cli to run locally
// TODO: import bookmarks via export html
// TODO: openai integration for tagging

func main() {
	if err := Run(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func Run() error {
	ctx := context.Background()

	config := config.NewConfig()
	if err := config.Load(".env"); err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	mode := config.Mode

	var logger *slog.Logger
	if mode == "prod" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	logger.Info(fmt.Sprintf("running in %s mode", mode))

	db, err := postgres.Connect(ctx, config.DBURL)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %v", err)
	}
	defer db.Close()

	server := server.NewServer()

	tasksStore := domaintasks.NewStore(db)
	tasksService := domaintasks.NewService(tasksStore)
	tasksHandler := apptasks.NewHandler(logger, tasksService)
	apptasks.AddRoutes(server, tasksHandler)

	diaryStore := domaindiary.NewStore(db)
	diaryService := domaindiary.NewService(diaryStore, tasksService)
	diaryHandler := appdiary.NewHandler(logger, diaryService)
	appdiary.AddRoutes(server, diaryHandler)

	logger.Info(fmt.Sprintf("starting server on port:%d", config.HostPort))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.HostPort), server); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	return nil
}
