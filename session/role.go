package session

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/just-arun/micro-session/model"
	"github.com/redis/go-redis/v9"
)

type roleAccess struct{}

func RoleAccess() roleAccess {
	return roleAccess{}
}

func (r roleAccess) Set(generalRedisDB *redis.Client, roleName string, data *model.Role) (err error) {
	ctx := context.Background()
	defer ctx.Done()
	key := fmt.Sprintf("role-%v", roleName)
	payload, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = generalRedisDB.Set(ctx, key, payload, time.Second*0).Err()
	return
}

func (r roleAccess) Get(generalRedisDB *redis.Client, roleName string) (data *model.Role, err error) {
	ctx := context.Background()
	defer ctx.Done()
	key := fmt.Sprintf("role-%v", roleName)

	result, err := generalRedisDB.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result), &data)
	return
}

func (r roleAccess) GetAccessesForRoles(generalRedisDB *redis.Client, roles []string) (accesses []string, err error) {
	ctx := context.Background()
	defer ctx.Done()
	var wg sync.WaitGroup
	wg.Add(len(roles))
	for _, role := range roles {
		roleFun := func() {
			res, err := r.Get(generalRedisDB, role)
			if err != nil {
				return
			}
			for _, v := range res.Accesses {
				accesses = append(accesses, v.Key)
			}
			wg.Done()
		}
		go roleFun()
	}
	wg.Wait()
	return
}
