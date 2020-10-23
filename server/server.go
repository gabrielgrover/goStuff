package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Coaster struct {
	Name         string
	Manufacturer string
	ID           string
	InPark       string
	Height       int
}

type coasterHandlers struct {
	sync.Mutex
	store map[string]Coaster
}

type adminPortal struct {
	password string
}

func main() {
	admin := newAdminPortal()
	coasterHandlers := newCoasterHandlers()
	http.HandleFunc("/coasters", coasterHandlers.coasters)
	http.HandleFunc("/coasters/", coasterHandlers.getCoaster)
	http.HandleFunc("/admin", admin.handler)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}
}

func newAdminPortal() *adminPortal {
	password := os.Getenv("ADMIN_PASSWORD")

	if password == "" {
		panic("required env var ADMIN_PASSWORD not set")
	}

	return &adminPortal{password: password}
}

func (a adminPortal) handler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()

	if !ok || user != "admin" || pass != a.password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - unauthorized"))
		return
	}

	w.Write([]byte("<html><h1>Super secret admin portal</h1></html>"))
}

func (h *coasterHandlers) coasters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	var coasters []Coaster

	h.Lock()

	for _, coaster := range h.store {
		coasters = append(coasters, coaster)
	}

	h.Unlock()

	jsonBytes, err := json.Marshal(coasters)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.Write(jsonBytes)
}

func (h *coasterHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")

	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("need content-type 'application/json'"))
		return
	}

	var coaster Coaster
	err = json.Unmarshal(bodyBytes, &coaster)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	coaster.ID = fmt.Sprintf("%d", time.Now().UnixNano())

	h.Lock()
	h.store[coaster.ID] = coaster
	defer h.Unlock()
}

func (h *coasterHandlers) getCoaster(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	coasters := h.store

	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	coasterID := parts[2]

	if coasterID == "random" {
		h.getRandomCoaster(w, r)
		return
	}

	h.Lock()
	coaster, ok := coasters[coasterID]
	h.Unlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(coaster)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.Write(jsonBytes)
}

func (h *coasterHandlers) getRandomCoaster(w http.ResponseWriter, r *http.Request) {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)
	targetIdx := randGen.Intn(len(h.store))

	i := 0

	h.Lock()
	defer h.Unlock()

	for _, coaster := range h.store {
		if i == targetIdx {
			jsonBytes, err := json.Marshal(coaster)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}

			w.Header().Add("content-type", "application/json")
			w.Write(jsonBytes)
		}

		i++
	}

	w.WriteHeader(http.StatusNotFound)
	return
}

func newCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{
			"id1": Coaster{
				Name:         "Fury 325",
				Height:       99,
				ID:           "id1",
				InPark:       "Carowinds",
				Manufacturer: "B+M",
			},
			"id2": Coaster{
				Name:         "Some name",
				Height:       101,
				ID:           "id2",
				InPark:       "Carowinds",
				Manufacturer: "B+M",
			},
		},
	}
}
