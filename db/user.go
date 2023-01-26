package db

import (
	"context"
	"fmt"
	"monitor/model"
	"time"
)

// gets a request and add the user (sign up)
func AddUser(s model.RegisterUserRequest) (int64, error) {
	ctx := context.Background()
	id := Rdb.Incr(ctx, "user_id")
	if id.Err() != nil {
		return 0, id.Err()
	}
	key := fmt.Sprintf("user:%d", id.Val())
	Rdb.HSet(ctx, key, model.User{
		UserName:  *s.UserName,
		PassWord:  *s.PassWord,
		CreatedAt: time.Now().Format(time.ANSIC),
	})
	return id.Val(), nil
}

func GetUser(id int64) (model.User, error) {
	key := fmt.Sprintf("user:%d", id)
	ctx := context.Background()
	u := model.User{}
	// get all the fields with the key
	res := Rdb.HGetAll(ctx, key)
	if res.Err() != nil {
		return u, res.Err()
	}
	// the user data
	u.UserName = res.Val()["user_name"]
	u.PassWord = res.Val()["password"]
	u.CreatedAt = res.Val()["created_at"]
	// return the user
	return u, nil
}
