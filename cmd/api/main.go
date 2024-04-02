package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"log/slog"

	"github.com/lewisje1991/code-bookmarks/internal/app/handlers"
	"github.com/lewisje1991/code-bookmarks/internal/app/router"
	"github.com/lewisje1991/code-bookmarks/internal/domain/bookmarks"
	"github.com/lewisje1991/code-bookmarks/internal/domain/notes"
	"github.com/lewisje1991/code-bookmarks/internal/platform/config"
	"github.com/lewisje1991/code-bookmarks/internal/platform/postgres"
	"github.com/lewisje1991/code-bookmarks/internal/platform/server"
)

// TODO: authentication
// TODO: authorization
// TODO: openai integration for tagging
// TODO: use expo to build a mobile app

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

	bookmarksStore := bookmarks.NewStore(db)
	booksmarksService := bookmarks.NewService(bookmarksStore)
	booksmarksHandler := handlers.NewBookmarkHandler(logger, booksmarksService)

	notesStore := notes.NewStore(db)
	notesService := notes.NewService(notesStore)
	notesHandler := handlers.NewNotesHandler(notesService, logger)

	server := server.NewServer()
	router.AddRoutes(server, booksmarksHandler, notesHandler)

	logger.Info(fmt.Sprintf("starting server on port:%d", config.HostPort))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.HostPort), server); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	return nil
}
