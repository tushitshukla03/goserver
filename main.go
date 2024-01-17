package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"video-call/server"
)

func main() {
	server.AllRooms.Init()
	go server.Broadcaster()

	router := mux.NewRouter()
	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprint(w,"hellow world")
	}).Methods("GET")

	// Your existing route handlers
	router.HandleFunc("/create", server.CreateRoomRequestHandler).Methods("GET")
	router.HandleFunc("/join", server.JoinRoomRequestHandler).Methods("GET")
	router.HandleFunc("/get", server.GetAllUser).Methods("GET")

	// Enable CORS for all routes
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	// Wrap the router with CORS middleware
	handler := c.Handler(router)

	log.Println("Starting Server on Port 8000")
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/",router)
}

