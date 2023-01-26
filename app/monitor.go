package app

import (
	"context"
	"monitor/config"
	"monitor/db"
	"monitor/model"
	"net/http"
	"strconv"
	"time"
)

// sends a request, stores the results and the database, and in case of failure, it increments the link's failures
func sendRequest(l model.Link, userId int64) model.Request {
	result := model.Request{}
	cli := &http.Client{
		Timeout: config.ClinetTimeOut,
	}
	req, _ := http.NewRequestWithContext(context.Background(),
		l.Method, l.Url, nil)
	response, _ := cli.Do(req)
	result.CreatedAt = time.Now().Format(time.ANSIC)
	result.Status = "failed"
	if response != nil && response.StatusCode < 300 && response.StatusCode >= 200 {
		result.Status = "success"
	}
	db.AddRequest(result, userId, l.LinkID)
	if result.Status == "failed" {
		db.IncreaseFailure(userId, l.LinkID)
	}
	return result
}

// collect all the links and send a request to each link
// Since our go routines are guranteed to exit after config.ClientTimeOut, we don't need to cancel the go routines
func updateUser(userID int64) {
	links, err := db.GetAllLink(userID)
	if err != nil || len(links) == 0 {
		return
	}
	for _, v := range links {
		go sendRequest(v, userID)
	}
}

// spawn a new go routine to update each user
func update() {
	ctx := context.Background()
	user_id := db.Rdb.Get(ctx, "user_id")
	// if our query fails, periodically call it again
	if user_id.Err() != nil {
		for {
			user_id = db.Rdb.Get(ctx, "user_id")
			if user_id.Err() == nil {
				break
			}
			<-time.After(time.Second * 1)
		}
	}
	n, _ := strconv.ParseInt(user_id.Val(), 10, 64)
	for i := int64(1); i <= n; i++ {
		go updateUser(i)
	}

}

// periodically calls the upadate function
func monitor() {
	for {
		update()
		<-time.After(config.WaitDuration)
	}
}
