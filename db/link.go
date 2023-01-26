package db

import (
	"context"
	"fmt"
	"monitor/config"
	Err "monitor/error"
	"monitor/model"
	"strconv"
	"time"
)

// Insert a verified link request object  into the database
// our key for the link is ["user:<user_id>:link<link_id>"]
func AddLink(l model.RegisterLinkRequest, id int64) (int64, error) {
	numberLinkKey := fmt.Sprintf("user:%d:link", id)
	ctx := context.Background()
	linkId := Rdb.Incr(ctx, numberLinkKey)
	if linkId.Err() != nil {
		return 0, linkId.Err()
	}
	// periodically tries to undo the incrementation
	if linkId.Val() > config.MaxNumLink {
		for {
			res := Rdb.Decr(ctx, numberLinkKey)
			if res.Err() == nil {
				break
			}
			<-time.After(time.Second * 1)
		}
		return 0, Err.ErrMaxNumLink
	}
	// Generate the user link key
	key := fmt.Sprintf("user:%d:link:%d", id, linkId.Val())
	Rdb.HSet(ctx, key, model.Link{
		LinkID:     linkId.Val(),
		Url:        *l.Url,
		ThreshHold: *l.ThreshHold,
		CreatedAt:  time.Now().Format(time.ANSIC),
		Failures:   0,
		Method:     *l.Method,
	})
	// returns the id of the generated link
	return linkId.Val(), nil
}

// gets a  link with the user and link id, returns an error in case of problem
func GetLink(userID int64, linkID int64) (model.Link, error) {
	key := fmt.Sprintf("user:%d:link:%d", userID, linkID)
	ctx := context.Background()
	l := model.Link{}
	res := Rdb.HGetAll(ctx, key)
	if res.Err() != nil {
		return l, res.Err()
	}
	l.LinkID = linkID
	l.Failures, _ = strconv.Atoi(res.Val()["failures"])
	l.CreatedAt = res.Val()["created_at"]
	l.Method = res.Val()["method"]
	l.Url = res.Val()["url"]
	l.ThreshHold, _ = strconv.Atoi(res.Val()["thresh_hold"])
	return l, nil
}

// gets all the links of a user with its id (uses the GetLink function)
func GetAllLink(userId int64) ([]model.Link, error) {
	list := make([]model.Link, 0)
	numberLinkKey := fmt.Sprintf("user:%d:link", userId)
	ctx := context.Background()
	linkId := Rdb.Get(ctx, numberLinkKey)
	if linkId.Err() != nil {
		return nil, linkId.Err()
	}
	n, _ := strconv.ParseInt(linkId.Val(), 10, 64)
	for i := int64(1); i <= n; i++ {
		l, err := GetLink(userId, i)
		if err == nil {
			list = append(list, l)
		}
	}
	return list, nil
}

// if a request fails, we increment its failure number
func IncreaseFailure(userId, link_id int64) {
	key := fmt.Sprintf("user:%d:link:%d", userId, link_id)
	ctx := context.Background()
	res := Rdb.HIncrBy(ctx, key, "failures", 1)
	// periodically increment in case of error
	if res.Err() != nil {
		for {
			res = Rdb.HIncrBy(ctx, key, "failures", 1)
			if res.Err() == nil {
				break
			}
			<-time.After(time.Second * 1)
		}
	}
}

// there will be no problem when incrementing or decrementing by multiple threads because redis is thread safe
