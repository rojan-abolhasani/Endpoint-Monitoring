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

// it adds the user user to the database (with AddUser in db)
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// The sign up method must be POST
	if r.Method != "POST" {
		ErrHandle(Err.ErrNotFound).ServeHTTP(w, r)
		return
	}
	s, err := util.ParseUserSignUp(r)
	if err != nil {
		ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
		return
	}
	// validate the User
	ok := util.ValidateUser(s)
	if !ok {
		ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
		return
	}
	// add the user
	id, err := db.AddUser(s)
	if err != nil {
		ErrHandle(Err.ErrInternal).ServeHTTP(w, r)
		return
	}
	// send the response if added successfully
	json.NewEncoder(w).Encode(model.RegisterUserResponse{
		Status: "success",
		UserId: id,
	})
}

// We gave a threshold for the maximum failures, send a warning if the limit is crossed
func UserWarningAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrHandle(Err.ErrNotFound).ServeHTTP(w, r)
		return
	}
	userId, _ := strconv.ParseInt(r.Header["User_id"][0], 10, 64)
	// get all the linkss
	list, err := db.GetAllLink(userId)
	if err != nil {
		ErrHandle(Err.ErrInternal).ServeHTTP(w, r)
		return
	}
	new_list := make([]model.Link, 0)
	// append every link that crosses the limit
	for _, v := range list {
		if v.Failures >= v.ThreshHold {
			new_list = append(new_list, v)
		}
	}
	// return a list from all the links that need the warning
	json.NewEncoder(w).Encode(new_list)
}
