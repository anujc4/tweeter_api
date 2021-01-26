package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
	"github.com/gorilla/schema"
)

// Set a Decoder instance as a package global, because it caches
// meta-data about structs, and an instance can be shared safely.
var decoder = schema.NewDecoder()

func (env *HttpApp) CreateTweet(w http.ResponseWriter, req *http.Request) {
	var request request.CreateUserRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	if err := request.ValidateCreateTweetRequest(); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	tweet, err := appModel.CreateTweet(&request)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSONwithStatus(w, http.StatusCreated, response.TransformUserResponse(*tweet))
}
