package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
	"github.com/gorilla/schema"
)


//type data_user struct {
//	status_code int `json:"status"`
//	Status string `json:"data"`
//	usrObj
//}

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
	params := mux.Vars(req)
	user_id := params["user_id"]
	id, err := strconv.Atoi(user_id)
	if err != nil {
		app.RenderErrorJSON(w, app.NewParseFormError(err))
	}
	fmt.Println(id)
	appModel := model.NewAppModel(req.Context(), env.DB)
	user, err := appModel.GetUserById(id)
	if err := req.ParseForm(); err != nil {
		app.RenderErrorJSON(w, app.NewParseFormError(err))
		return
	}
	fmt.Println(user)
	app.RenderJSON(w, response.TransformUserResponse(*user))

}

func (env *HttpApp) UpdateUser(w http.ResponseWriter, req *http.Request) {
	var request request.UpdateUserRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		app.RenderJSONwithStatus(w, http.StatusBadRequest, "Error -Invalid Request Body")
		return
	}
	params := mux.Vars(req)
	id1:= params["user_id"]
	id, err1 := strconv.Atoi(id1)
	if err1 != nil {
		app.RenderJSONwithStatus(w, http.StatusBadRequest, "Error -Invalid User Id")
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	user,err := appModel.UpdateUser(&request,id)
	if err != nil {
		fmt.Println("Error",err)
		app.RenderJSONwithStatus(w, http.StatusBadRequest, "Error-Bad Request")
		return
	}
	app.RenderJSONwithStatus(w,http.StatusOK, response.TransformUserResponse(*user))
	return
}

func (env *HttpApp) DeleteUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	user_id := params["user_id"]
	id, err := strconv.Atoi(user_id)
	if err != nil {
		app.RenderJSONwithStatus(w, http.StatusBadRequest, "Error -Invalid User Id")
	}
	fmt.Println(id)
	appModel := model.NewAppModel(req.Context(), env.DB)
	user, err1 := appModel.DeleteUser(id)
	if err1 != nil {
		fmt.Println("Error",err1)
		app.RenderJSONwithStatus(w, http.StatusBadRequest, "Error-Invalid User Id")
		return
	}
	fmt.Println(user)
	app.RenderJSON(w, response.TransformUserResponse(*user))
}
