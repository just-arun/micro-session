package session

import (
	"context"
	"encoding/json"

	"github.com/just-arun/micro-session/model"
	"github.com/redis/go-redis/v9"
)

type siteMap struct{}

func SiteMap() siteMap {
	return siteMap{}
}

func (st siteMap) Set(generalRedisDB *redis.Client, data []model.ServiceMap) (err error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return
	}
	return generalRedisDB.Set(context.Background(), "site-map", payload, 0).Err()
}

func (st siteMap) Get(generalRedisDB *redis.Client) (data *model.ServiceMap, err error) {
	result, err := generalRedisDB.Get(context.Background(), "site-map").Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result), &data)
	return
}
