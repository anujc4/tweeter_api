package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
	"github.com/gorilla/mux"
)

// Set a Decoder instance as a package global, because it caches
// meta-data about structs, and an instance can be shared safely.

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
	app.RenderJSONwithStatus(w, http.StatusCreated, response.TransformTweetResponse(*tweet))
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
	tweets, err := appModel.GetTweets(&request)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	resp := response.MapTweetResponse(*tweets, response.TransformTweetResponse)
	app.RenderJSON(w, resp)
}

func (env *HttpApp) GetTweetByID(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement thismake
	params := mux.Vars(req)
	idstring := params["tweet_id"]

	id, err := strconv.Atoi(idstring)

	if err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	tweet, errs := appModel.GetTweetById(id)

	if err != nil {
		app.RenderErrorJSON(w, errs)
		return
	}

	app.RenderJSONwithStatus(w, http.StatusCreated, response.TransformTweetResponse(*tweet))

}

func (env *HttpApp) UpdateTweet(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idstring := params["tweet_id"]

	id, err := strconv.Atoi(idstring)

	if err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	if err := req.ParseForm(); err != nil {
		{
			app.RenderErrorJSON(w, app.NewParseFormError(err).SetCode(http.StatusBadRequest))
			return
		}
	}

	var request request.CreateTweetRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)

	tweet, errs := appModel.UpdateTweetById(&request, id)

	if errs != nil {
		app.RenderErrorJSON(w, errs)
		return
	}

	app.RenderJSONwithStatus(w, http.StatusCreated, response.TransformTweetResponse(*tweet))

}

func (env *HttpApp) DeleteTweet(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idstring := params["tweet_id"]

	id, err := strconv.Atoi(idstring)

	if err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	errs := appModel.DeleteTweetById(id)

	if err != nil {
		app.RenderErrorJSON(w, errs)
		return
	}

	app.RenderJSONwithStatus(w, http.StatusCreated, "tweet deleted Successfully")

}
