package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
)

func (env *HttpApp) CreateTweet(w http.ResponseWriter, r *http.Request) {
	var req request.CreateTweetRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	if err := req.ValidateCreateTweetRequest(); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	claims, ok := r.Context().Value("claims").(*request.UserSessionClaims)
	if !ok {
		err := app.NewError(errors.New("parse_error")).SetCode(http.StatusInternalServerError)
		app.RenderErrorJSON(w, err)
		return
	}

	appModel := model.NewAppModel(r.Context(), env.DB)
	tweet, err := appModel.CreateTweet(claims.UserID, req.Content, req.ParentTweet)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSONwithStatus(w, http.StatusCreated, tweet)
}
