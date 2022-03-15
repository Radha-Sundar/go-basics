package db

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)
