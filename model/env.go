package model

import "time"

type redisEnv struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type otp struct {
	ExpireTime time.Duration `mapstructure:"expireTime"`
}

type natsEnv struct {
	Token string `mapstructure:"token"`
}

type Env struct {
	UserSession    redisEnv `mapstructure:"userSession"`
	GeneralSession redisEnv `mapstructure:"generalSession"`
	Secret         string   `mapstructure:"secret"`
	OTP            otp      `mapstructure:"otp"`
	Nats           natsEnv  `mapstructure:"nats"`
}
