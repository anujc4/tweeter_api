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
	"github.com/gorilla/schema"
)

// Set a Decoder instance as a package global, because it caches
// meta-data about structs, and an instance can be shared safely.
var decoder = schema.NewDecoder()

//CreateUser method in handler user.go
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

//GetUsers method in handler user.go
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

//GetUsersByID method in handler user.go
func (env *HttpApp) GetUserByID(w http.ResponseWriter, req *http.Request) {
	userID, err := getID(req)

	if err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	var request request.GetUserByIDRequest

	if decodeErr := decoder.Decode(&request, req.Form); decodeErr != nil {
		app.RenderErrorJSON(w, app.NewError(decodeErr).SetCode(http.StatusBadRequest))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)

	user, getUserErr := appModel.GetUserByID(userID, &request)
	if getUserErr != nil {
		app.RenderErrorJSON(w, getUserErr)
		return
	}

	app.RenderJSONwithStatus(w, http.StatusCreated, response.TransformUserResponse(*user))
}

//UpdateUser method in handler user.go
func (env *HttpApp) UpdateUser(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement this

	userID, err := getID(req)
	if err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	var request request.UpdateUserRequest
	decoder := json.NewDecoder(req.Body)
	if error := decoder.Decode(&request); error != nil {
		app.RenderErrorJSON(w, app.NewError(error))
		return
	}

	if validateErr := request.ValidateUpdateUserRequest(); validateErr != nil {
		app.RenderErrorJSON(w, app.NewError(validateErr))
		return
	}
	appModel := model.NewAppModel(req.Context(), env.DB)
	user, modelErr := appModel.UpdateUser(userID, &request)
	if modelErr != nil {
		app.RenderErrorJSON(w, modelErr)
		return
	}
	app.RenderJSONwithStatus(w, http.StatusCreated, response.TransformUserResponse(*user))

}

//DeleteUser method in handler user.go
func (env *HttpApp) DeleteUser(w http.ResponseWriter, req *http.Request) {
	userID, err := getID(req)
	if err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}
	appModel := model.NewAppModel(req.Context(), env.DB)
	deleteErr := appModel.DeleteUser(userID)
	if deleteErr != nil {
		app.RenderErrorJSON(w, deleteErr)
		return
	}

	// fix response  "User got deleted" maybe??
	app.RenderJSON(w, "Not yet implemented!")
}

func getID(req *http.Request) (uint, error) {
	params := mux.Vars(req)
	userID, err := strconv.Atoi(params["user_id"])
	if err != nil {
		return uint(userID), err
	}
	return uint(userID), nil
}
