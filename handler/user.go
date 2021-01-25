package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

// Set a Decoder instance as a package global, because it caches
// meta-data about structs, and an instance can be shared safely.
var decoder = schema.NewDecoder()

//CreateUser creates a new user
func (env *HttpApp) CreateUser(w http.ResponseWriter, req *http.Request) {
	var request request.CreateUserRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	if err := request.ValidateCreateUserRequest(); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	user, err := appModel.CreateUser(&request)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSONwithStatus(w, http.StatusCreated, response.TransformUserResponse(*user))
}

// GetUsers returns all users
func (env *HttpApp) GetUsers(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		app.RenderErrorJSON(w, app.NewParseFormError(err))
		return
	}

	var request request.GetUsersRequest
	if err := decoder.Decode(&request, req.Form); err != nil {
		app.RenderErrorJSON(w, app.NewError(err).SetCode(http.StatusBadRequest))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	users, err := appModel.GetUsers(&request)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	resp := response.MapUsersResponse(*users, response.TransformUserResponse)
	app.RenderJSON(w, resp)
}

// GetUserByID returns a user by given user id
func (env *HttpApp) GetUserByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userID := vars["user_id"]

	appModel := model.NewAppModel(req.Context(), env.DB)
	user, err := appModel.GetUserByID(userID)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	resp := response.TransformUserResponse(user)
	app.RenderJSON(w, resp)
}

// UpdateUser updates a user details in db
func (env *HttpApp) UpdateUser(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement this
	app.RenderJSON(w, "Not yet implemented!")
}

//DeleteUser delets a user from db
func (env *HttpApp) DeleteUser(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement this
	app.RenderJSON(w, "Not yet implemented!")
}
