# Monitoring the Endpoints in http

This project is written in go with RESTful API.

Note : tun the project using `go run main.go`

## Database

####User
| UserName  | PasswWord | CreatedAt | 
| :-------- | --------- | ----------|
| string    | string    | string    |

####Link
| LinkID  | Url       | ThreshHold | CreatedAt | Method | Failure
| :------ | --------- | -----------| --------- | -------| -------
| int64   | string    | int        | string    | string | int
