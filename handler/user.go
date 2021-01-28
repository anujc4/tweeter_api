package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
	"github.com/gorilla/schema"
	"github.com/gorilla/mux"
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
	// TODO: Implement this
	//app.RenderJSON(w, "Not yet implemented!")
	params := mux.Vars(req)
	userid := params["user_id"]
	id,err := strconv.Atoi(userid)

	if err != nil {
		app.RenderErrorJSON(w,app.NewError(err))
		return 
	}
	parseErr := req.ParseForm()
	if parseErr != nil {
		app.RenderErrorJSON(w,app.NewParseFormError(parseErr))
		return 
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	users, getUserErr := appModel.GetUserByID(id)
	if getUserErr != nil {
		app.RenderErrorJSON(w,getUserErr)
		return 
	}

	resp := response.TransformUserResponse(*users)
	app.RenderJSON(w, resp)
}

func (env *HttpApp) UpdateUser(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement this
	//app.RenderJSON(w, "Not yet implemented!")
	var request request.CreateUserRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}
	appModel := model.NewAppModel(req.Context(), env.DB)

	params := mux.Vars(req)
	userid:= params["user_id"]
	id, err1 := strconv.Atoi(userid)
		if err1 != nil {
				app.RenderErrorJSON(w, app.NewParseFormError(err1))
		}
		if err := request.ValidateCreateUserRequest(); err != nil {
	 		app.RenderErrorJSON(w, app.NewError(err))
	 		return
	 	}
	err := appModel.UpdateUser(&request,id)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSON(w, "updated user ")
}

func (env *HttpApp) DeleteUser(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement this
	//app.RenderJSON(w, "Not yet implemented!")
	params := mux.Vars(req)
	userid:= params["user_id"]
	id, err1 := strconv.Atoi(userid)
		if err1 != nil {
			app.RenderErrorJSON(w, app.NewParseFormError(err1))
		}
	if err := req.ParseForm(); err != nil {
		app.RenderErrorJSON(w, app.NewParseFormError(err))
		return
	}
	appModel := model.NewAppModel(req.Context(), env.DB)

	err := appModel.DeleteUser(id)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSON(w, "deleted user ")
}
