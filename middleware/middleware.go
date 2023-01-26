package middleware

import (
	"monitor/config"
	Err "monitor/error"
	"monitor/view"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
)

// sets the return type to Json
// Http handler gets request and response and writes in the response
func SetContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// response header
		head := w.Header()
		// set the type of the response
		head.Add("content-type", "application/json")
		// calls the next handler
		h.ServeHTTP(w, r)
	})
}

// check if the type of the request is Json
func CheckContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if the content type is not "application/json" return an error
		if r.Header.Get("content-type") != "application/json" {
			view.ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
			return
		}
		// calls the next handler
		h.ServeHTTP(w, r)
	})
}

// uses jwt to authorize the user (with token)
func Authorization(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			view.ErrHandle(Err.ErrNoPermission).ServeHTTP(w, r)
			return
		}
		token, err := jwt.Parse(r.Header["Authorization"][0], func(t *jwt.Token) (interface{}, error) {
			return config.SecretKey, nil
		})
		// return an error when the token is not valid
		if err != nil || !token.Valid {
			view.ErrHandle(Err.ErrNoPermission).ServeHTTP(w, r)
			return
		}
		// claims is the token payload (in the request)
		claims, ok := token.Claims.(jwt.MapClaims)
		// If sth is not right
		if !ok {
			view.ErrHandle(Err.ErrNoPermission).ServeHTTP(w, r)
			return
		}
		// gets the id and puts the user_id for the next handlers
		user_id := int64(claims["id"].(float64))
		r.Header.Add("User_id", strconv.FormatInt(user_id, 10))
		// calls the next handler
		h.ServeHTTP(w, r)
	})
}
