package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anujc4/tweeter_api/handler"
	"github.com/anujc4/tweeter_api/internal/env"
	"github.com/anujc4/tweeter_api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func routeHandler(env *env.Env) {
	h := handler.NewHttpApp(env)
	mw := middleware.NewMiddlewareApp(env)

	router := mux.NewRouter().StrictSlash(true)

	apiV1 := router.PathPrefix("/v1").Subrouter()
	apiV1.Use(mw.Authentication)

	// Health Check Endpoint
	router.HandleFunc("/simple_health", h.SimpleHealthCheck).Methods("GET")
	router.HandleFunc("/detail_health", h.DetailedHealthCheck).Methods("GET")

	router.HandleFunc("/signup", h.CreateUser).Methods("POST")
	router.HandleFunc("/login", h.Login).Methods("POST")
	apiV1.HandleFunc("/logout", h.Logout).Methods("GET")

	// Users API
	apiV1.HandleFunc("/users", h.GetUsers).Methods("GET")
	apiV1.HandleFunc("/user/{user_id}", h.GetUserByID).Methods("GET")
	apiV1.HandleFunc("/user/{user_id}", h.UpdateUser).Methods("PUT")
	apiV1.HandleFunc("/user/{user_id}", h.DeleteUser).Methods("DELETE")

	// Tweets
	// TODO: Implement the h
	// apiV1.HandleFunc("/tweet", h.CreateTweet).Methods("POST")
	// apiV1.HandleFunc("/tweets", h.GetTweets).Methods("GET")
	// apiV1.HandleFunc("/tweet/{tweet_id}", h.GetTweetByID).Methods("GET")
	// apiV1.HandleFunc("/tweet/{tweet_id}", h.UpdateTweet).Methods("PUT")
	// apiV1.HandleFunc("/tweet/{tweet_id}", h.UpdateTweet).Methods("DELETE")

	// Start the server
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	// log.Fatal(http.ListenAndServe(":3000", handlers.RecoveryHandler()(loggedRouter)))
	log.Fatal(http.ListenAndServe(":3000", loggedRouter))
}

func main() {
	env := env.Init()

	fmt.Println("Starting Tweeter API on port 3000...")

	routeHandler(env)
}
