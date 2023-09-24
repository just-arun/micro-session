package session

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/just-arun/micro-session/model"
	"github.com/redis/go-redis/v9"
)

type otp struct{}

func OTP() otp {
	return otp{}
}

func (st otp) SetOTP(rDB *redis.Conn, otp model.OTP, expireOne time.Duration) error {
	ctx := context.Background()
	defer ctx.Done()
	key := fmt.Sprintf("%v-%v-%v", otp.Key, otp.UserID, otp.OTP)
	payload, err := json.Marshal(otp)
	if err != nil {
		return err
	}
	return rDB.Set(ctx, key, payload, expireOne).Err()
}

func (st otp) GetOTP(rDB *redis.Conn, otp model.OTP, expireOne time.Duration) (data *model.OTP, err error) {
	ctx := context.Background()
	defer ctx.Done()
	key := fmt.Sprintf("%v-%v-%v", otp.Key, otp.UserID, otp.OTP)
	result, err := rDB.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result), &data)
	return
}

func (st otp) GetAndDelOTP(rDB *redis.Conn, otp model.OTP, expireOne time.Duration) (data *model.OTP, err error) {
	ctx := context.Background()
	defer ctx.Done()
	key := fmt.Sprintf("%v-%v-%v", otp.Key, otp.UserID, otp.UserID)
	result, err := rDB.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result), &data)
	return
}
