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
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	// public routes
	s.Router.Post("/login", handlers.ToHttpHandlerFunc(h.LoginHandler))
	s.Router.Post("/signup", handlers.ToHttpHandlerFunc(h.SignupHandler))

	// private routes
	s.Router.Group(func(r chi.Router) {
		r.Use(handlers.ToMiddleware(h.AuthMiddlewareFunc))

		r.Get("/me", handlers.ToHttpHandlerFunc(h.GetUserHandler))

		r.Route("/bookmark", func(r chi.Router) {
			r.Post("/", handlers.ToHttpHandlerFunc(h.CreateBookmarkHandler))

			r.Route("/{id}", func(r chi.Router) {
				r.Use(handlers.ToMiddleware(h.BookmarkMiddlewareFunc))

				r.Get("/", handlers.ToHttpHandlerFunc(h.GetBookmarkHandler))
				r.Patch("/", handlers.ToHttpHandlerFunc(h.UpdateBookmarkHandler))
				r.Delete("/", handlers.ToHttpHandlerFunc(h.DeleteBookmarkHandler))
				// router.Post("/{id}/summarise", ToHttpHandlerFunc(s.SummariseBookmarkHandler))
			})
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
