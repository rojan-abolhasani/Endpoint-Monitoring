package view

import (
	"encoding/json"
	"monitor/auth"
	"monitor/config"
	"monitor/db"
	Err "monitor/error"
	"monitor/model"
	"monitor/util"
	"net/http"
	"time"
)

// returns the Token after generation in the authorixation
func GetToken(w http.ResponseWriter, r *http.Request) {
	// the mothed must be GET
	if r.Method != "GET" {
		ErrHandle(Err.ErrNotFound).ServeHTTP(w, r)
		return
	}
	s, err := util.ParseToken(r)
	if err != nil {
		ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
		return
	}
	u, err := db.GetUser(*s.UserId)
	if err != nil {
		ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
		return
	}
	// check the password: if true, return the token
	if u.PassWord != *s.PassWord {
		ErrHandle(Err.ErrNoPermission).ServeHTTP(w, r)
		return
	}
	// create the token
	token, err := auth.CreateToken(*s.UserId)
	if err != nil {
		ErrHandle(Err.ErrInternal).ServeHTTP(w, r)
		return
	}
	// return the reponse containing the Token
	json.NewEncoder(w).Encode(model.TokenResponse{
		Status:  "success",
		ExpDate: time.Now().Add(config.TokenDuration).Format(time.ANSIC),
		Token:   token,
	})

}
