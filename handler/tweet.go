package handler

import (
	"encoding/json"
	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"fmt"
)

//var decoder = schema.NewDecoder()

func (env *HttpApp) CreateTweet(w http.ResponseWriter, req *http.Request)  {
	var request request.CreateTweetRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil{
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	if err := request.ValidateCreateTweetRequest(); err != nil{
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	tweet, err := appModel.CreateTweet(&request)
	if err != nil{
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSONwithStatus(w, http.StatusCreated, response.TransformTweetResponse(*tweet))
}

func (env *HttpApp) GetTweets(w http.ResponseWriter, req *http.Request)  {
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
func (env *HttpApp) GetTweetByID(w http.ResponseWriter, req *http.Request)  {
	params := mux.Vars(req)
	tweet_id := params["tweet_id"]
	id, err := strconv.Atoi(tweet_id)
	if err != nil {
		app.RenderErrorJSON(w, app.NewParseFormError(err))
	}
	fmt.Println(id)
	appModel := model.NewAppModel(req.Context(), env.DB)
	tweet, err := appModel.GetTweetById(id)
	if err := req.ParseForm(); err != nil {
		app.RenderErrorJSON(w, app.NewParseFormError(err))
		return
	}
	fmt.Println(tweet)
	app.RenderJSON(w, response.TransformTweetResponse(*tweet))

}
func (env *HttpApp) UpdateTweet(w http.ResponseWriter, req *http.Request)  {
	var request request.UpdateTweetRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil{
		app.RenderJSONwithStatus(w, http.StatusBadRequest, "Error -Invalid Request Body")
		return
	}
	params := mux.Vars(req)
	id1:= params["tweet_id"]
	id, err1 := strconv.Atoi(id1)
	if err1 != nil {
		app.RenderJSONwithStatus(w, http.StatusBadRequest, "Error -Invalid Tweet Id")
	}
	appModel := model.NewAppModel(req.Context(), env.DB)
	tweet, err := appModel.UpdateTweet(&request, id)
	if err != nil {
		fmt.Println("Error",err)
		app.RenderJSONwithStatus(w, http.StatusBadRequest, "Error-Bad Request")
		return
	}
	app.RenderJSONwithStatus(w,http.StatusOK, response.TransformTweetResponse(*tweet))
	return
}

func (env *HttpApp) DeleteTweet(w http.ResponseWriter, req *http.Request)  {
	params := mux.Vars(req)
	tweet_id := params["tweet_id"]
	id, err := strconv.Atoi(tweet_id)
	if err != nil {
		app.RenderJSONwithStatus(w, http.StatusBadRequest, "Error -Invalid Tweet Id")
	}
	fmt.Println(id)
	appModel := model.NewAppModel(req.Context(), env.DB)
	tweet, err1 := appModel.DeleteTweet(id)
	if err1 != nil {
		fmt.Println("Error",err1)
		app.RenderJSONwithStatus(w, http.StatusBadRequest, "Error-Invalid Tweet Id")
		return
	}
	fmt.Println(tweet)
	app.RenderJSON(w, response.TransformTweetResponse(*tweet))
}

