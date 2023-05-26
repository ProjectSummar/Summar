package routes

import (
	"fmt"
	"log"
	"net/http"
	"summar/server/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	Address  string
	Handlers *handlers.Handlers
}

func NewRouter(address string, handlers *handlers.Handlers) *Router {
	return &Router{
		Address:  address,
		Handlers: handlers,
	}
}

func (r *Router) Run() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// public routes
	router.Post("/login", handlers.ToHttpHandlerFunc(r.Handlers.LoginHandler))
	router.Post("/signup", handlers.ToHttpHandlerFunc(r.Handlers.SignupHandler))

	// private routes
	router.Group(func(router chi.Router) {
		router.Use(handlers.ToMiddleware(r.Handlers.AuthMiddlewareFunc))
		router.Get("/me", handlers.ToHttpHandlerFunc(r.Handlers.GetUserHandler))
		router.Route("/bookmark", func(router chi.Router) {
			router.Post("/create", handlers.ToHttpHandlerFunc(r.Handlers.CreateBookmarkHandler))
			// router.Get("/get", handlers.ToHttpHandlerFunc(r.Handlers.GetBookmarkHandler))
			// router.Post("/update", ToHttpHandlerFunc(s.UpdateBookmarkHandler))
			// router.Post("/delete", ToHttpHandlerFunc(s.DeleteBookmarkHandler))
			// router.Post("/summarise", ToHttpHandlerFunc(s.SummariseBookmarkHandler))
		})
	})

	chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})
	log.Println("Server running on port", r.Address)
	log.Fatal(http.ListenAndServe(":"+r.Address, router))
}
