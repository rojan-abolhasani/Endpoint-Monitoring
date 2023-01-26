package model

// these are the structures that we expect to be sent to us in the request, otherwise it faces an Error

type TokenRequest struct {
	UserId   *int64  `json:"user_id"`
	PassWord *string `json:"password"`
}

type RegisterUserRequest struct {
	UserName *string `json:"user_name"`
	PassWord *string `json:"password"`
}

type RegisterLinkRequest struct {
	Url        *string `json:"url"`
	ThreshHold *int    `json:"thresh_hold"`
	Method     *string `json:"method"`
}

type LinkRequest struct {
	LinkId *int64 `json:"link_id"`
}
