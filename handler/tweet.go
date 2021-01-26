package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
	"github.com/gorilla/mux"
)

//CreateTweet is endpoint to create a tweet -> /tweet [POST]
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

// GetTweets is api endpoint to get all tweets -> /tweets [GET]
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

// GetTweetByID is api endpoint to get tweet by id -> /tweet/{tweet_id} [GET]
func (env *HttpApp) GetTweetByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	tweetID := vars["tweet_id"]

	appModel := model.NewAppModel(req.Context(), env.DB)
	tweet, err := appModel.GetTweetByID(&tweetID)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	resp := response.TransformTweetResponse(*tweet)
	app.RenderJSON(w, resp)
}

/*
// UpdateTweet is api endpoint to update a tweet -> /tweet/{tweet_id} [PUT]
func (env *HttpApp) UpdateTweet(w http.ResponseWriter, req *http.Request) {
	var request request.UpdateUserRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	if err := request.ValidateUpdateUserRequest(); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	vars := mux.Vars(req)
	userID := vars["user_id"]
	request.ID = userID

	appModel := model.NewAppModel(req.Context(), env.DB)
	user, err := appModel.UpdateUser(request)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	resp := response.TransformUserResponse(user)
	app.RenderJSON(w, resp)
}
*/

//DeleteTweet is api endpoint to delete a tweet -> /user/{tweet_id} [DELETE]
func (env *HttpApp) DeleteTweet(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	TweetID := vars["tweet_id"]

	appModel := model.NewAppModel(req.Context(), env.DB)
	err := appModel.DeleteTweet(&TweetID)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSONwithStatus(w, 204, "")
}
