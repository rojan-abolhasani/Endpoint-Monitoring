package router

import (
	Err "monitor/error"
	"monitor/middleware"
	"monitor/view"
	"net/http"
)

// view shows the response
// creates our router binding our endpoints to our handlers
func Router() *http.ServeMux {
	mux := http.NewServeMux()
	setType := middleware.SetContentType
	checkType := middleware.CheckContentType
	authurize := middleware.Authorization
	mux.Handle("/api/link", setType(checkType(authurize(http.HandlerFunc(view.FetchLink)))))
	mux.Handle("/api/warnings", setType(checkType(authurize(http.HandlerFunc(view.UserWarningAll)))))
	mux.Handle("/api/links", setType(checkType(authurize(http.HandlerFunc(view.FetchAllLink)))))
	mux.Handle("/api/token", setType(checkType(http.HandlerFunc(view.GetToken))))
	mux.Handle("/api/register/link", setType(checkType(authurize(http.HandlerFunc(view.RegisterLink)))))
	mux.Handle("/api/register/user", setType(checkType(http.HandlerFunc(view.RegisterUser))))
	mux.Handle("/", setType(view.ErrHandle(Err.ErrNotFound)))
	return mux
}
