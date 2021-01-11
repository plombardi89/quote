// Copyright 2019 Philip Lombardi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	"github.com/plombardi89/gozeug/randomzeug"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var port = 8080

const (
	EnvPORT        = "PORT"
	EnvHOST        = "HOST"
	EnvOpenAPIPath = "OPENAPI_PATH"
)

type Server struct {
	id       string
	host     string
	port     int
	router   *chi.Mux
	upgrader websocket.Upgrader
	hub      *Hub
	random   *randomzeug.Random
	quotes   []string
}

type QuoteResult struct {
	Server string    `json:"server"`
	Quote  string    `json:"quote"`
	Time   time.Time `json:"time"`
}

func (s *Server) GetQuote(w http.ResponseWriter, r *http.Request) {
	quote := s.random.RandomSelectionFromStringSlice(s.quotes)
	res := QuoteResult{
		Server: s.id,
		Quote:  quote,
		Time:   time.Now().UTC(),
	}

	resJson, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")

	if _, err := w.Write(resJson); err != nil {
		log.Panicln(err)
	}
}

func (s *Server) StreamQuotes(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{hub: s.hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.readPump()
	go client.writePump()
}

func (s *Server) ConfigureRouter() {
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)

	s.router.Get("/", s.GetQuote)
	s.router.HandleFunc("/ws", s.StreamQuotes)

	s.router.Get(getEnv(EnvOpenAPIPath, "/.well-known/openapi.yaml"), s.GetOpenAPIDocument)
}

func (s *Server) Start() error {
	s.hub = newHub(s.random, s.quotes)
	go s.hub.run()

	listenAddr := fmt.Sprintf("%s:%d", s.host, s.port)
	log.Printf("listening on %s\n", listenAddr)
	return http.ListenAndServe(listenAddr, s.router)
}

func main() {
	if portString := os.Getenv(EnvPORT); portString != "" {
		p, err := strconv.Atoi(portString)
		if err != nil {
			log.Fatalln(err)
		}

		if p < 1 || p > 65535 {
			log.Fatalln("Server port must be in range 1..65535 (inclusive)")
		}

		port = p
	}

	startingQuotes := []string{
		"Abstraction is ever present.",
		"A late night does not make any sense.",
		"A principal idea is omnipresent, much like candy.",
		"Nihilism gambles with lives, happiness, and even destiny itself!",
		"The light at the end of the tunnel is interdependent on the relatedness of motivation, subcultures, and management.",
		"Utter nonsense is a storyteller without equal.",
		"Non-locality is the driver of truth. By summoning, we vibrate.",
		"A small mercy is nothing at all?",
		"The last sentence you read is often sensible nonsense.",
		"668: The Neighbor of the Beast.",
	}

	random := randomzeug.NewRandom()
	s := Server{
		id:     getServerID(os.Hostname, random),
		host:   os.Getenv(EnvHOST),
		port:   port,
		router: chi.NewRouter(),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		random: random,
		quotes: startingQuotes,
	}

	s.ConfigureRouter()

	log.Fatalln(s.Start())
}
