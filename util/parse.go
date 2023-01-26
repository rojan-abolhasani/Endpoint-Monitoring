package util

import (
	"encoding/json"
	Err "monitor/error"
	"monitor/model"
	"net/http"
)

// (used in view), reads the body of the request and checks if it has the same structure as we expect, otherwise it returns an error
func ParseUserSignUp(r *http.Request) (model.RegisterUserRequest, error) {
	s := model.RegisterUserRequest{}
	json.NewDecoder(r.Body).Decode(&s)
	if s.PassWord == nil || s.UserName == nil {
		return s, Err.ErrMissingFields
	}
	return s, nil
}

// same is ParseUserSignUp but for the tokens
func ParseToken(r *http.Request) (model.TokenRequest, error) {
	s := model.TokenRequest{}
	json.NewDecoder(r.Body).Decode(&s)
	if s.UserId == nil || s.PassWord == nil {
		return s, Err.ErrMissingFields
	}
	return s, nil
}

// same is ParseUserSignUp but for adding links
func ParseRegisterLink(r *http.Request) (model.RegisterLinkRequest, error) {
	l := model.RegisterLinkRequest{}
	json.NewDecoder(r.Body).Decode(&l)
	if l.ThreshHold == nil || l.Url == nil || l.Method == nil || !ValidateMethod(*l.Method) {
		return l, Err.ErrMissingFields
	}
	return l, nil
}

// same is ParseUserSignUp but for getting links
func ParseLinkGet(r *http.Request) (int64, error) {
	l := model.LinkRequest{}
	json.NewDecoder(r.Body).Decode(&l)
	if l.LinkId == nil {
		return 0, Err.ErrMissingFields
	}
	return *l.LinkId, nil
}
