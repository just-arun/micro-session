package model

import "github.com/redis/go-redis/v9"

type GlobalCtx struct {
	GeneralSessionRedisDB *redis.Client
	UserSessionRedisDB    *redis.Client
	Env                   *Env
}
