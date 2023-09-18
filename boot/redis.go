package boot

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Redis(address, password string, db int, description ...interface{}) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	defer ctx.Done()
	fmt.Println(fmt.Sprintf("%v redis connected...", description...))
	return rdb
}
