package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
	// "github.com/gorilla/schema"
	"github.com/gorilla/mux"

)

// var decoder = schema.NewDecoder()

//create tweet

func (env *HttpApp) CreateTweet(w http.ResponseWriter, req *http.Request) {
	var request request.CreateTweetRequest
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
	resp := response.MapTweetsResponse(*tweets, response.TransformTweetResponse)
	app.RenderJSON(w, resp)
}

func (env *HttpApp) GetTweetByID(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement this
	params := mux.Vars(req)
	id1:= params["tweet_id"]
	id, err1 := strconv.Atoi(id1)
    if err1 != nil {
      app.RenderErrorJSON(w, app.NewParseFormError(err1))
    }

	if err := req.ParseForm(); err != nil {
		app.RenderErrorJSON(w, app.NewParseFormError(err))
		return
	}
	appModel := model.NewAppModel(req.Context(), env.DB)
	tweet, err := appModel.GetTweetByID(id)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	resp := response.TransformTweetResponse(*tweet)
	app.RenderJSON(w, resp)

	// app.RenderJSON(w, "Not yet implemented!")
}

func (env *HttpApp) UpdateTweet(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement this
	var request request.CreateTweetRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	params := mux.Vars(req)
	id1:= params["tweet_id"]
	id, err1 := strconv.Atoi(id1)
		if err1 != nil {
				app.RenderErrorJSON(w, app.NewError(err1))
		}
    if err := request.ValidateCreateTweetRequest(); err != nil {
      app.RenderErrorJSON(w, app.NewError(err))
      return
    }



	appModel := model.NewAppModel(req.Context(), env.DB)
	err := appModel.UpdateTweet(&request,id)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSON(w, "updated")
	// app.RenderJSON(w, "Not yet implemented!")
}

func (env *HttpApp) DeleteTweet(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement this
	params := mux.Vars(req)
	id1:= params["tweet_id"]
	id, err1 := strconv.Atoi(id1)
		if err1 != nil {
					app.RenderErrorJSON(w, app.NewParseFormError(err1))
		}
	if err := req.ParseForm(); err != nil {
		app.RenderErrorJSON(w, app.NewParseFormError(err))
		return
	}
	appModel := model.NewAppModel(req.Context(), env.DB)
	err := appModel.DeleteTweet(id)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSON(w, "deleted")
	// app.RenderJSON(w, "Not yet implemented!")
}
