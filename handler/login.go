package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/internal/constants"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (env *HttpApp) Login(w http.ResponseWriter, r *http.Request) {
	var req request.LoginRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	if err := req.ValidateLoginRequest(); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(r.Context(), env.DB)
	user, err := appModel.VerifyUserCredential(&req)
	if err != nil {
		e := err.
			SetCode(http.StatusUnauthorized).
			SetMessage("User not found. Please check your email and try again")
		app.RenderErrorJSON(w, e)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		e := app.NewError(err).
			SetCode(http.StatusUnauthorized).
			SetMessage("Invalid Credentials. Please try again")
		app.RenderErrorJSON(w, e)
		return
	}

	id := uuid.NewString()

	claims := request.UserSessionClaims{
		UserID: user.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().AddDate(0, 0, 3).Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
			Id:        id,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err1 := token.SignedString(env.PrivateKey)
	if err1 != nil {
		e := app.NewError(err1).
			SetCode(http.StatusInternalServerError).
			SetMessage("Unable to create user session")
		app.RenderErrorJSON(w, e)
		return
	}

	env.Redis.SAdd(r.Context(), fmt.Sprintf("%s%d", constants.RedisJwtPrefix, user.ID), id)

	resp := response.LoginResponse{
		Email: user.Email,
		Token: signedToken,
	}

	app.RenderJSON(w, resp)
}

func (env *HttpApp) Logout(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*request.UserSessionClaims)
	if !ok {
		e := app.NewError(errors.New("parsing error")).
			SetCode(http.StatusInternalServerError).
			SetMessage("Unable to process request.")
		app.RenderErrorJSON(w, e)
	}

	// Invaludate the token
	env.Redis.SRem(r.Context(), fmt.Sprintf("%s%d", constants.RedisJwtPrefix, claims.UserID), claims.Id)

	w.WriteHeader(http.StatusNoContent)
}
