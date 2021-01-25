package middleware

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/internal/env"
	"github.com/anujc4/tweeter_api/request"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type middlewareApp struct {
	DB        *gorm.DB
	PublicKey *rsa.PublicKey
}

func NewMiddlewareApp(env *env.Env) *middlewareApp {
	return &middlewareApp{
		DB:        env.DB,
		PublicKey: env.PublicKey,
	}
}

func (mw *middlewareApp) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("YAHAN AAYA")
		authenticationHeader := r.Header.Get("Authorization")
		if authenticationHeader == "" {
			e := errors.New("missing authorization token")
			appError := app.
				NewError(e).
				SetCode(http.StatusUnauthorized).
				SetMessage("Missing session details. Please login and try again.")
			app.RenderErrorJSON(w, appError)
			return
		}

		authHeaderData := strings.Split(authenticationHeader, "Bearer ")
		if len(authHeaderData) != 2 {
			e := errors.New("malformed token")
			appError := app.
				NewError(e).
				SetCode(http.StatusUnauthorized).
				SetMessage("Malformed token error. Please login and try again.")
			app.RenderErrorJSON(w, appError)
			return
		} else {
			tokenString := authHeaderData[1]

			token, _ := jwt.ParseWithClaims(tokenString,
				&request.UserSessionClaims{},
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
						return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
					}
					return mw.PublicKey, nil
				})

			if claims, ok := token.Claims.(*request.UserSessionClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "claims", claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				appError := app.
					NewError(errors.New("unauthorised")).
					SetCode(http.StatusUnauthorized).
					SetMessage("User unauthorised. Please login and try again.")
				app.RenderErrorJSON(w, appError)
			}
		}
	})
}
