package config

import "time"

// MySQL配置
type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Db       string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

// jwt配置
type JwtConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

// Redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
	Timeout  int    `mapstructure:"timeout"`
}

// app元数据
type MyAppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
	Local   string `mapstructure:"local"`
}

// Nacos配置
//type NacosConfig struct {
//	Host      string `mapstructure:"host"`
//	Port      uint64 `mapstructure:"port"`
//	Namespace string `mapstructure:"namespace"`
//	User      string `mapstructure:"user"`
//	Password  string `mapstructure:"password"`
//	DataId    string `mapstructure:"dataid"`
//	Group     string `mapstructure:"group"`
//}

// 日志配置
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// snowflake 配置
type SnowConfig struct {
	StartTime time.Time `mapstructure:"start_time"`
	MachineID int64     `mapstructure:"machine_id"`
}
type ServerConfig struct {
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	//NacosInfo NacosConfig `mapstructure:"nacos" json:"nacos"`
	JWTInfo   JwtConfig   `mapstructure:"jwt" json:"jwt"`
	MyAppInfo MyAppConfig `mapstructure:"app" json:"app"`
	LogInfo   LogConfig   `mapstructure:"log" json:"log"`
	RedisInfo RedisConfig `mapstructure:"redis" json:"redis"`
	SnowInfo  SnowConfig  `mapstructure:"snowflake" json:"snowflake"`
}
