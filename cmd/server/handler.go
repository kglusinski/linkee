package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"linkee/internal/linkee"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMuxHandler() *mux.Router {
	r := mux.NewRouter()
	makeHandlers(r)

	return r
}

func makeHandlers(r *mux.Router) {
	r.HandleFunc("/", home)
	r.HandleFunc("/{page-slug}", getPage)
	r.HandleFunc("/{page-slug}/{link-slug}", updateCounter).Methods("PUT", "PATCH")
}

func home(w http.ResponseWriter, r *http.Request) {

}

func getPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageSlug := vars["page-slug"]
	log.Info().
		Str("pageSlug", pageSlug).
		Str("method", "getPage").
		Send()

	db := connect()
	repo := linkee.NewMySQLRepository(db)
	svc := linkee.NewService(repo)

	page := svc.GetPage(pageSlug)

	res, err := json.Marshal(page)
	if err != nil {
		log.Err(err).Msg("Couldn't marshal page")
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Err(err).Msg("Couldn't write response")
	}
}

func updateCounter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageSlug := vars["page-slug"]
	linkSlug := vars["link-slug"]
	log.Info().
		Str("pageId", pageSlug).
		Str("linkId", linkSlug).
		Str("method", "updateCounter").
		Send()
	db := connect()
	repo := linkee.NewMySQLRepository(db)
	svc := linkee.NewService(repo)

	svc.UpdateCounter(pageSlug, linkSlug)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func connect() *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("root:test@tcp(localhost:%d)/linkee_db", 33062))
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Couldn't connect to the database")
	}

	return db
}
