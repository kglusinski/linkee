package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"linkee/internal/linkee"
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
	r.HandleFunc("/{pageId}/{linkId}", updateCounter)
}

func home(w http.ResponseWriter, r *http.Request) {

}

func getPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageId := vars["pageId"]
	log.Info().
		Str("pageId", pageId).
		Str("method", "getPage").
		Send()

	repo := linkee.NewInMemoryRepository()
	svc := linkee.NewService(repo)

	page := svc.GetPage(pageId)

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
	pageId := vars["pageId"]
	linkId := vars["linkId"]
	log.Info().
		Str("pageId", pageId).
		Str("linkId", linkId).
		Str("method", "updateCounter").
		Send()
	repo := linkee.NewInMemoryRepository()
	svc := linkee.NewService(repo)

	svc.UpdateCounter(pageId, linkId)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}