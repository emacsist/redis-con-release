package config

type AppConfig struct {
	Comment string     `json:"comment"`
	Redis   []RedisEle `json:"redis"`
	Idle    int        `json:"idle"`
}

type RedisEle struct {
	Host   string `json:"host"`
	Passwd string `json:"passwd"`
	Port   int    `json:"port"`
}
