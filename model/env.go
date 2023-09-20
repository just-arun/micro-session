package model

type Redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Env struct {
	UserSession    Redis  `mapstructure:"userSession"`
	GeneralSession Redis  `mapstructure:"generalSession"`
	Secret         string `mapstructure:"secret"`
}
