package common

import "github.com/micro/go-micro/v2/config"

// MysqlConfig mysql相关配置
type MysqlConfig struct {
	Host string `json:"host"`
	Port int64 `json:"port"`
	User string `json:"user"`
	Pwd string `json:"pwd"`
	Database string `json:"database"`
}

// GetMysqlConfigFromConsul 获取 mysql配置
func GetMysqlConfigFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}
