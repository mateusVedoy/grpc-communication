package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	RenderChi "github.com/go-chi/render"
)

type Router struct {
	controller *Controller
}

func (R *Router) Start() {
	contentType := middleware.AllowContentType("application/json")
	route := chi.NewRouter()
	route.Use(middleware.RequestID)
	route.Use(middleware.RealIP)
	route.Use(middleware.Recoverer)
	route.Use(contentType)
	route.Use(RenderChi.SetContentType(RenderChi.ContentTypeJSON))
	route.Use(middleware.Timeout(60 * time.Second))

	route.Post("/message/create", R.controller.CreateMessage)

	panic(http.ListenAndServe(":8080", route))
}

func NewRouter() *Router {
	return &Router{
		controller: NewController(),
	}
}
