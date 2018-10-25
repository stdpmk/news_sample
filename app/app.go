package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/stdpmk/news_sample/db"
)

type App struct {
	db     *db.DB
	router *mux.Router
}

func NewApp(opts *db.ConnOpts) *App {

	d := db.NewDatabase(opts)
	app = &App{
		db: d,
	}
	return app

}

func (a *App) Close() {
	a.db.Close()
}

func (a *App) addRoute(path string, f http.HandlerFunc) *mux.Route {
	return a.router.HandleFunc(path, f)
}

// Глобальный инстанс приложения
// Будет доступен из http хендлеров
var app *App

func (a *App) setupMiddleware() http.Handler {

	handler := handlers.RecoveryHandler(
		handlers.PrintRecoveryStack(true),
	)(a.router)
	handler = handlers.LoggingHandler(os.Stdout, handler)
	handler = SetJsonTypeHandler(handler)
	handler = cors.AllowAll().Handler(handler)

	return handler
}

func (a *App) setupRouters() http.Handler {
	m := mux.NewRouter().Schemes("http").PathPrefix("/v1").Subrouter()
	a.router = m

	// news
	a.addRoute("/news", CreateNews).Methods("POST")
	a.addRoute("/news", GetNewsList).Methods("GET")
	a.addRoute("/news/{id}", GetNews).Methods("GET")
	a.addRoute("/news/{id}", UpdateNews).Methods("PUT")

	// comments
	a.addRoute("/comments", CreateComment).Methods("POST")
	a.addRoute("/comments/{id}", GetComment).Methods("GET")
	a.addRoute("/comments", GetCommentList).Methods("GET")
	a.addRoute("/comments/{id}", UpdateComment).Methods("PUT")
	a.addRoute("/comments/{id}", DeleteComment).Methods("DELETE")

	// author
	a.addRoute("/authors", CreateAuthor).Methods("POST")
	a.addRoute("/authors", GetAuthorList).Methods("GET")
	a.addRoute("/authors/{id}", GetAuthor).Methods("GET")
	a.addRoute("/authors/{id}", UpdateAuthor).Methods("PUT")
	a.addRoute("/authors/{id}", DeleteAuthor).Methods("DELETE")

	handler := a.setupMiddleware()
	return handler
}

func (a *App) Run(port string) error {

	handler := a.setupRouters()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: handler,
	}
	var err error
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return err
	}
	return nil
}

// Middleware

func SetJsonTypeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, req)
	})
}
