# Monitoring the Endpoints in http

This project is written in go with RESTful API.

Note : run the project using `go run main.go`

## Database

This database is a key-value database, so id cannot be seen in the tables.

##### User:
| UserName  | PasswWord | CreatedAt | 
| :-------- | --------- | ----------|
| string    | string    | string    |

(Auto-increment by user_id)

##### Link:
| LinkID  | Url       | ThreshHold | CreatedAt | Method | Failure
| :------ | --------- | -----------| --------- | -------| -------
| int64   | string    | int        | string    | string | int

##### Request:
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

**Add Link:**

`POST /api/register/link`
(Authorization with token in header)

request structure 

```
{
	"url":"https://google.com",
	"thresh_hold": 5 , // just an exmaple
  	"method" : GET // only 3 methods were allowed
}
```

**Get Link:**

`GET /api/link`
(Authorization with token in header)

request structure 

```
{
	"link_id":2 // exmaple
}
```

**Get All Links:**

`GET /api/links`
(Authorization with token in header)

No body is needed here


**Get Token:**

`GET /api/token `

request structure 

```
{
	"user_id":2, //example
	"password": "*bar*" 
}	
```

**Warnings:**

`GET /api/warnings`
(Authorization with token in header)

No body is needed here


#### Responses:

##### Error Response:
```
{
	status : string 
	error_msg : string 
	help : string 
}	
```

##### Sign-Up Response:
```
{
	status : string 
	user_id : int64 
}	
```

##### Add Link Response:
```
{
	status : string 
	link_id : int64 
}	
```

##### Token Response:
```
{
	status : string 
	exp_date : string 
	token : string 
}	
```

##### Add Link Response:
```
{
	Link
	requests []Request 
}	
```






