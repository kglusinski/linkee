package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

func NewMuxHandler() *mux.Router {
	r := mux.NewRouter()
	makeHandlers(r)

	return r
}

func makeHandlers(r *mux.Router) {
	r.HandleFunc("/", home)
	r.HandleFunc("/{pageId}", getPage)
}

func home(w http.ResponseWriter, r *http.Request) {

}

func getPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageId := vars["pageId"]

	log.Info().Msg(pageId)
}