// config/config.go
package config

import (
	"strconv"
	"time"
)

// DBConfig 数据库配置
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	Timeout  time.Duration
}

// PortStr 返回端口的字符串形式
func (c *DBConfig) PortStr() string {
	return strconv.Itoa(c.Port)
}

// GetDBConfig 返回数据库配置
func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "127.0.0.1", // MySQL 服务器地址
		Port:     3307,        // 端口
		User:     "root",      // 用户名 (替换为您的)
		Password: "cssao888",  // 密码 (替换为您的)
		Name:     "todo_db",   // 数据库名
		SSLMode:  "false",
		Timeout:  10 * time.Second,
	}
}
