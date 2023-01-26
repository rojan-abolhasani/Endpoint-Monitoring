package config

import (
	"time"

	"github.com/redis/go-redis/v9"
)

// secret key for signing it should be loaded from a file or env variable and
// be generated from a cryptographically secure hash function
var SecretKey = []byte("Anything really!")

// our database configs it should be loaded from a file or env variable
const addr string = "127.0.0.1:6379"
const password = ""
const db int = 0

// our redis database option
var DbOptions = redis.Options{
	Addr:     addr,
	Password: password,
	DB:       db,
}

// our api reference for people to refer to
const Help = "anthing"

// link config
const MaxNumLink = 20

// fetch config
const ClinetTimeOut = time.Second * 30

// server config
const Addr = "127.0.0.1:8080"
const ReadTimeOut = time.Second * 30
const WriteTimeOut = time.Second * 30
const IdleTimeOut = time.Second * 30

//update config

const WaitDuration = time.Minute * 10

const TokenDuration = time.Hour * 2
