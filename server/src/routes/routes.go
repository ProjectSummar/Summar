package routes

import (
	"log"
	"net/http"
	"summar/server/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	router := chi.NewRouter()

	return &Server{
		Router: router,
	}
}

func (s *Server) MountHandlers(h *handlers.Handlers) {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// public routes
	router.Post("/login", handlers.ToHttpHandlerFunc(h.LoginHandler))
	router.Post("/signup", handlers.ToHttpHandlerFunc(h.SignupHandler))

	// private routes
	router.Group(func(router chi.Router) {
		router.Use(handlers.ToMiddleware(h.AuthMiddlewareFunc))
		router.Get("/me", handlers.ToHttpHandlerFunc(h.GetUserHandler))
		router.Route("/bookmark", func(router chi.Router) {
			router.Post("/", handlers.ToHttpHandlerFunc(h.CreateBookmarkHandler))
			router.Get("/{id}", handlers.ToHttpHandlerFunc(h.GetBookmarkHandler))
			router.Patch("/{id}", handlers.ToHttpHandlerFunc(h.UpdateBookmarkHandler))
			router.Delete("/{id}", handlers.ToHttpHandlerFunc(h.DeleteBookmarkHandler))
			// router.Post("/summarise", ToHttpHandlerFunc(s.SummariseBookmarkHandler))
		})
	})
}

func (s *Server) Run(address string) {
	chi.Walk(s.Router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})

	log.Println("Server running on port", address)
	log.Fatal(http.ListenAndServe(":"+address, s.Router))
}
