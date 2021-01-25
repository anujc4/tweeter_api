package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
	"github.com/dgrijalva/jwt-go"
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

	claims := request.UserSessionClaims{
		UserID: user.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().AddDate(0, 0, 3).Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
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

	resp := response.LoginResponse{
		Email: user.Email,
		Token: signedToken,
	}

	app.RenderJSON(w, resp)
}
