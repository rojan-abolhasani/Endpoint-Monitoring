package view

import (
	"encoding/json"
	"monitor/db"
	Err "monitor/error"
	"monitor/model"
	"monitor/util"
	"net/http"
	"strconv"
)

func RegisterLink(w http.ResponseWriter, r *http.Request) {
	// check the method (it should be POST for sign up)
	if r.Method != "POST" {
		ErrHandle(Err.ErrNotFound).ServeHTTP(w, r)
		return
	}
	// returns the request object (in util)
	l, err := util.ParseRegisterLink(r)
	if err != nil {
		ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
		return
	}
	// reads the user_id from the request header (it was added in the Authorization (in middleware))
	user_id, _ := strconv.ParseInt(r.Header.Get("User_id"), 10, 64)
	// add the link
	id, err := db.AddLink(l, user_id)
	if err != nil {
		ErrHandle(Err.ErrInternal).ServeHTTP(w, r)
		return
	}
	// writes the response
	json.NewEncoder(w).Encode(model.RegisterLinkResponse{
		Status: "success",
		LinkId: id,
	})
}

// same as the FetchLink (it's for getting all the links)
func FetchAllLink(w http.ResponseWriter, r *http.Request) {
	// the method must be GET
	if r.Method != "GET" {
		ErrHandle(Err.ErrNotFound).ServeHTTP(w, r)
		return
	}
	userId, _ := strconv.ParseInt(r.Header["User_id"][0], 10, 64)
	// save the links in the list
	list, err := db.GetAllLink(userId)
	if err != nil {
		ErrHandle(Err.ErrInternal).ServeHTTP(w, r)
		return
	}
	// view the links
	json.NewEncoder(w).Encode(list)
}

// gets a link with today's requests
func FetchLink(w http.ResponseWriter, r *http.Request) {
	// method should be get
	if r.Method != "GET" {
		ErrHandle(Err.ErrNotFound).ServeHTTP(w, r)
		return
	}
	linkId, err := util.ParseLinkGet(r)
	if err != nil {
		ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
		return
	}
	userId, _ := strconv.ParseInt(r.Header["User_id"][0], 10, 64)
	link, err := db.GetLink(userId, linkId)

	if err != nil {
		ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
		return
	}
	// save the request
	requests, err := db.GetTodayRequest(userId, linkId)
	if err != nil {
		ErrHandle(Err.ErrInternal).ServeHTTP(w, r)
		return
	}
	// response contains a link and the requests' list
	res := model.LinkResponse{
		Link:     link,
		Requests: requests,
	}
	// view the response
	json.NewEncoder(w).Encode(res)
}
