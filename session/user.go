package session

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/just-arun/micro-session/model"
	"github.com/redis/go-redis/v9"
)

type userSession struct{}

func UserSession() userSession {
	return userSession{}
}

func (r userSession) Set(sessionRedisDB *redis.Client, sessionID string, userID uint) (err error) {
	ctx := context.Background()
	defer ctx.Done()
	key := fmt.Sprintf("session-%v-%v", sessionID, userID)
	payload := model.UserSessionData{UserID: userID, Time: time.Now()}
	da, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = sessionRedisDB.Set(ctx, key, da, time.Second*0).Err()
	return
}

func (r userSession) GetOneBySessionID(sessionRedisDB *redis.Client, sessionID string) (data *model.UserSessionData, err error) {
	ctx := context.Background()
	defer ctx.Done()
	keys := fmt.Sprintf("session-%v-*", sessionID)
	val, err := sessionRedisDB.Keys(ctx, keys).Result()
	if err != nil {
		return
	}
	if len(val) < 1 {
		return
	}
	key := val[0]
	da, err := sessionRedisDB.Get(ctx, key).Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(da), &data)
	return
}

func (r userSession) GetManyByUserID(sessionRedisDB *redis.Client, userID uint) (data []model.UserSessionData, err error) {
	ctx := context.Background()
	defer ctx.Done()
	keys := fmt.Sprintf("session-*-%v", userID)
	val, err := sessionRedisDB.Keys(ctx, keys).Result()
	if err != nil {
		return
	}
	if len(val) < 1 {
		return
	}
	count := len(val)

	c := make(chan model.UserSessionData, count)
	defer close(c)

	result := func() {
		for _, k := range val {
			go r.GetSessionByKey(sessionRedisDB, k, c)
		}
	}
	
	result()

	for i := 0; i < count; i++ {
		data = append(data, <-c)
	}

	return
}

func (r userSession) GetSessionByKey(sessionRedisDB *redis.Client, key string, c chan model.UserSessionData) (err error) {
	ctx := context.Background()
	defer ctx.Done()
	da, err := sessionRedisDB.Get(ctx, key).Result()
	var data model.UserSessionData
	err = json.Unmarshal([]byte(da), &data)
	if err != nil {
		return err
	}
	c <- data
	return
}
