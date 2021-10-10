package server

import (
	"embed"
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"
	"go.uber.org/zap"
)

// Config defines the configuration for the longform server.
type Config struct {
	ListenAddr string `default:":8080" help:"port to listen on"`
}

// Server defines the necessary components to host the longform server.
type Server struct {
	log      *zap.Logger
	cfg      Config
	staticFS embed.FS
}

// New creates a new longform server based on the provided config and static directory.
func New(log *zap.Logger, cfg Config, staticFS embed.FS) (s *Server, err error) {
	s = &Server{
		log:      log,
		cfg:      cfg,
		staticFS: staticFS,
	}

	return s, nil
}

// Close closes the server.
func (server *Server) Close() error {
	return nil
}

// Serve starts the longform server.
func (server *Server) Serve() error {
	// construct a standard go http handler
	handler := http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			server.logDebug(req)

			if req.URL.Path == "/" {
				w.Write([]byte("hello world"))
				return
			}
			if served := server.ServeStatic(w, req); served {
				return
			}

			w.WriteHeader(http.StatusNotFound)
		})

	errsCh := make(chan error, 1)
	go func() {
		server.log.Info("Starting server", zap.String("listen addr", server.cfg.ListenAddr))
		errsCh <- http.ListenAndServe(server.cfg.ListenAddr, handler)
	}()
	return <-errsCh
}

func (server *Server) logDebug(req *http.Request) {
	server.log.Debug("request",
		zap.String("host", req.Host),
		zap.String("hostname", req.URL.Hostname()),
		zap.String("url", req.URL.String()),
		zap.String("referer", req.Referer()),
		zap.String("user agent", req.UserAgent()),
		zap.String("forwarded for", req.Header.Get("X-Forwarded-For")))
}

// serveJSONError writes JSON error to response output stream.
func (server *Server) serveJSONError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	server.log.Debug("sending json error to client", zap.Error(err))

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		server.log.Error("failed to write json error response", zap.Error(errs.Wrap(err)))
	}
}
