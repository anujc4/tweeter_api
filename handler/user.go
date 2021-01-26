package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

func (env *HttpApp) GetUserByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	user_id, parse_err := strconv.Atoi(vars["user_id"])
	if user_id == 0 || parse_err != nil {
		err := app.NewValidationError(errors.New("Invalid user_id"))
		app.RenderErrorJSON(w, err)
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	user, err := appModel.GetUser(map[string]interface{}{"id": user_id})
	if err != nil {
		err = err.
			SetCode(http.StatusNotFound).
			SetMessage(fmt.Sprintf("No user found with id %d", user_id))
		app.RenderErrorJSON(w, err)
		return
	}
	resp := response.TransformUserResponse(*user)
	app.RenderJSON(w, resp)
}

func (env *HttpApp) UpdateUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	user_id, parse_err := strconv.ParseUint(vars["user_id"], 10, 0)
	if user_id == 0 || parse_err != nil {
		err := app.NewValidationError(errors.New("Invalid user_id"))
		app.RenderErrorJSON(w, err)
		return
	}

	var request request.UpdateUserRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	user, err := appModel.UpdateUser(uint(user_id), &request)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSON(w, response.TransformUserResponse(*user))
}

func (env *HttpApp) DeleteUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	user_id, parse_err := strconv.ParseUint(vars["user_id"], 10, 0)
	if user_id == 0 || parse_err != nil {
		err := app.NewValidationError(errors.New("Invalid user_id"))
		app.RenderErrorJSON(w, err)
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	rows, err := appModel.DeleteUser(uint(user_id))
	if err != nil {
		e := app.
			NewError(err).
			SetCode(http.StatusInternalServerError).
			SetMessage(fmt.Sprintf("Unable to delete user with id %d", user_id))
		app.RenderErrorJSON(w, e)
		return
	}

	if rows == 0 {
		e := app.
			NewError(errors.New("invalid user_id")).
			SetCode(http.StatusBadRequest).
			SetMessage(fmt.Sprintf("No user found with id %d", user_id))
		app.RenderErrorJSON(w, e)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
