package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
)

func (env *HttpApp) CreateTweet(w http.ResponseWriter, req *http.Request) {
	var request request.CreateTweetRequest
	decoder := json.NewDecoder(req.Body)
	if decodeErr := decoder.Decode(&request); decodeErr != nil {
		app.RenderErrorJSON(w, app.NewError(decodeErr))
		return
	}

	if validateErr := request.ValidateCreateTweetRequest(); validateErr != nil {
		app.RenderErrorJSON(w, app.NewError(validateErr))
		return
	}
	appModel := model.NewAppModel(req.Context(), env.DB)
	tweet, creationErr := appModel.CreateTweet(&request)
	if creationErr != nil {
		app.RenderErrorJSON(w, creationErr)
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
