package model

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type UserSessionType redis.Client
type GeneralSessionType redis.Client

type UserSessionData struct {
	UserID uint
	Time   time.Time
}
