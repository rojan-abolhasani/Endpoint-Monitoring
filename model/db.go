package model

// how we save User in the database
type User struct {
	UserName  string `redis:"user_name"`
	PassWord  string `redis:"password"`
	CreatedAt string `redis:"created_at"`
}

// how we save link in the database
type Link struct {
	LinkID     int64  `redis:"link_id"`
	Url        string `redis:"url"`
	ThreshHold int    `redis:"thresh_hold"`
	CreatedAt  string `redis:"created_at"`
	Method     string `redis:"method"`
	Failures   int    `redis:"failures"`
}

// how we save request in the database
type Request struct {
	Status    string `redis:"status"`
	CreatedAt string `redis:"created_at"`
}
