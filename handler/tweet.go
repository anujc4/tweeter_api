package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/gorilla/schema"
)

// Set a Decoder instance as a package global, because it caches
// meta-data about structs, and an instance can be shared safely.
var decoderr = schema.NewDecoder()

func (env *HttpApp) CreateTweet(w http.ResponseWriter, req *http.Request) {
	var request request.CreateTweetRequest
	decoderr := json.NewDecoder(req.Body)
	if err := decoderr.Decode(&request); err != nil {
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

// func (env *HttpApp) GetUsers(w http.ResponseWriter, req *http.Request) {
// 	if err := req.ParseForm(); err != nil {
// 		app.RenderErrorJSON(w, app.NewParseFormError(err))
// 		return
// 	}
// 	var request request.GetUsersRequest
// 	if err := decoder.Decode(&request, req.Form); err != nil {
// 		app.RenderErrorJSON(w, app.NewError(err).SetCode(http.StatusBadRequest))
// 		return
// 	}
// 	appModel := model.NewAppModel(req.Context(), env.DB)
// 	users, err := appModel.GetUsers(&request)
// 	if err != nil {
// 		app.RenderErrorJSON(w, err)
// 		return
// 	}
// 	resp := response.MapUsersResponse(*users, response.TransformUserResponse)
// 	app.RenderJSON(w, resp)
// }

// func (env *HttpApp) GetUserByID(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	id := params["user_id"]
// 	if err := req.ParseForm(); err != nil {
// 		app.RenderErrorJSON(w, app.NewParseFormError(err))
// 		return
// 	}
// 	var request request.GetUsersRequest
// 	if err := decoder.Decode(&request, req.Form); err != nil {
// 		app.RenderErrorJSON(w, app.NewError(err))
// 		return
// 	}

// 	appModel := model.NewAppModel(req.Context(), env.DB)
// 	user, err := appModel.GetUserByID(&request, id)
// 	// fmt.Println(user.UpdatedAt)
// 	if err != nil {
// 		app.RenderErrorJSON(w, err)
// 		return
// 	}
// 	resp := response.TransformUserResponse(*user)
// 	app.RenderJSON(w, resp)
// }

// func (env *HttpApp) UpdateUser(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	id := params["user_id"]
// 	fmt.Println(id)
// 	var request request.CreateUserRequest
// 	decoder := json.NewDecoder(req.Body)
// 	if err := decoder.Decode(&request); err != nil {
// 		app.RenderErrorJSON(w, app.NewError(err))
// 		return
// 	}

// 	appModel := model.NewAppModel(req.Context(), env.DB)
// 	user, err := appModel.UpdateUser(&request, id)
// 	if err != nil {
// 		app.RenderErrorJSON(w, err)
// 		return
// 	}
// 	app.RenderJSONwithStatus(w, http.StatusCreated, response.TransformUserResponse(*user))
// }

// func (env *HttpApp) DeleteUser(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	id := params["user_id"]
// 	if err := req.ParseForm(); err != nil {
// 		app.RenderErrorJSON(w, app.NewParseFormError(err))
// 		return
// 	}
// 	var request request.CreateUserRequest
// 	if err := decoder.Decode(&request, req.Form); err != nil {
// 		app.RenderErrorJSON(w, app.NewError(err))
// 		return
// 	}

// 	appModel := model.NewAppModel(req.Context(), env.DB)
// 	user, err := appModel.DeleteUserByID(&request, id)
// 	// fmt.Println(user.UpdatedAt)
// 	if err != nil {
// 		app.RenderErrorJSON(w, err)
// 		return
// 	}
// 	resp := response.TransformUserResponse(*user)
// 	app.RenderJSON(w, resp)
// }
