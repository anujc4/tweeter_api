package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
	"github.com/gorilla/mux"
)

func (env *HttpApp) CreateTweet(w http.ResponseWriter, req *http.Request) {
	var request request.CreateTweetRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	tweet, err := appModel.CreateTweet(&request)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSONwithStatus(w, http.StatusCreated, *tweet)
}

func (env *HttpApp) GetTweets(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		app.RenderErrorJSON(w, app.NewParseFormError(err))
		return
	}
	var request request.GetTweetsRequest
	if err := decoder.Decode(&request, req.Form); err != nil {
		app.RenderErrorJSON(w, app.NewError(err).SetCode(http.StatusBadRequest))
		return
	}
	appModel := model.NewAppModel(req.Context(), env.DB)
	users, err := appModel.GetTweets(&request)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	resp := response.MapTweetResponse(*users, response.TransformTweetResponse)
	app.RenderJSON(w, resp)
}

func (env *HttpApp) GetTweetByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["tweet_id"]
	if err := req.ParseForm(); err != nil {
		app.RenderErrorJSON(w, app.NewParseFormError(err))
		return
	}
	var request request.GetTweetsRequest
	if err := decoder.Decode(&request, req.Form); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	user, err := appModel.GetTweetById(&request, id)
	// fmt.Println(user.UpdatedAt)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	resp := response.TransformTweetResponse(*user)
	app.RenderJSON(w, resp)
}

func (env *HttpApp) UpdateTweet(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["tweet_id"]
	fmt.Println(id)
	var request request.CreateTweetRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	user, err := appModel.UpdateTweet(&request, id)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSONwithStatus(w, http.StatusCreated, response.TransformTweetResponse(*user))
}

func (env *HttpApp) DeleteTweet(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["tweet_id"]
	if err := req.ParseForm(); err != nil {
		app.RenderErrorJSON(w, app.NewParseFormError(err))
		return
	}
	var request request.CreateTweetRequest
	if err := decoder.Decode(&request, req.Form); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	user, err := appModel.DeleteTweetByID(&request, id)
	// fmt.Println(user.UpdatedAt)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	resp := response.TransformTweetResponse(*user)
	app.RenderJSON(w, resp)
}
