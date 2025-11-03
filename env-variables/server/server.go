package server

import (
	"fmt"
	"github.com/Disaxy/fast-track/tree/main/env-variables/config"
	"html"
	"log/slog"

	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("New request", "addr", r.RemoteAddr, "path", r.URL.Path)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Serve(config *config.Config) {
	http.HandleFunc("/index", indexHandler)

	s := &http.Server{
		Addr: config.Address(),
	}

	slog.Info("Server listening on " + config.Address())

	if err := s.ListenAndServe(); err != nil {
		slog.Error("Server failed", "error", err)
	}
}
