package handler

import (
	"fmt"
	"net/http"

	"github.com/anujc4/tweeter_api/internal/app"
)

type data struct {
	DB     string `json:"status"`
	Status string `json:"message"`
}

// SimpleHealthCheck responds with a "UP!" message every time it is invoked
// It is useful in scenarios where you just want to check if the server is
// running or not.
// In realtime, this API misht fail if case of two reasons:
// 1. Server not running due to some error
// 2. Taking too much time since server is under heacy load
func (env *HttpApp) SimpleHealthCheck(w http.ResponseWriter, r *http.Request) {
	app.RenderJSON(w, "UP!")
}

// DetailedHealthCheck responds with a detailed feedback of each system's
// status. Usually you would want to add your checks for DBs (if multiple),
// Redis, ElasticSearch or any other dependencies.
func (env *HttpApp) DetailedHealthCheck(w http.ResponseWriter, r *http.Request) {
	db, _ := env.DB.DB()

	if err := db.Ping(); err != nil {
		_ = fmt.Errorf("error reaching DB %s", err)
		err := app.NewError(err).
			SetCode(http.StatusInternalServerError).
			SetMessage("offline")

		app.RenderErrorJSON(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := data{
		DB:     "ONLINE",
		Status: "UP!",
	}

	app.RenderJSON(w, response)
}
