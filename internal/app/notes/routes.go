package notes

import (
	"github.com/lewisje1991/code-bookmarks/internal/platform/server"
)

func AddRoutes(server *server.Server, h *Handler) {
	// Notes
	server.AddRoute("POST", "/note", h.PostHandler())
}
