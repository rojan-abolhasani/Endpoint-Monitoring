# Monitoring the Endpoints in http

This project is written in go with RESTful API.

Note : tun the project using `go run main.go`

## Database

This database is a key-value database, so id cannot be seen in the tables.

####User
| UserName  | PasswWord | CreatedAt | 
| :-------- | --------- | ----------|
| string    | string    | string    |

(Auto-increment by user_id)

####Link
| LinkID  | Url       | ThreshHold | CreatedAt | Method | Failure
| :------ | --------- | -----------| --------- | -------| -------
| int64   | string    | int        | string    | string | int

####Request
| status    | CreatedAt | 
| :-------- | --------- |
| string    | string    | 


### Specs:

For all requests and responses we have `Content-Type: application/json`.

Authorization is with JWT.

#### User endpoints:

**Sign Up:**

`POST /api/register/user`

request structure 

```
{
	"user_name":"foo", // length >= 5
	"password":"*bar*" // alpha numeric with one character, length >= 10
}
```

**Adding Link:**

`POST /api/register/link`

request structure 

```
{
	"url":"https://google.com",
	"thresh_hold": 5 , // just an exmaple
  "method" : GET // only 3 methods were allowed
}
```



