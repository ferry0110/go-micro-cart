package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)



// GetConsulConfig 设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error){
	consulSource := consul.NewSource(
		//设置配置中心地址 host:port
		consul.WithAddress(host+":"+ strconv.FormatInt(port, 10)),
		//设置配置中心前缀，默认前缀为/micro/config
		consul.WithPrefix(prefix),
		//设置是否不带前缀,设为true表示 不带前缀也可以获得配置
		consul.StripPrefix(true),
		)
	curConfig, err := config.NewConfig()
	if err != nil  {
		return curConfig,err
	}
	//加载配置
	err = curConfig.Load(consulSource)
	if err != nil  {
		return curConfig,err
	}
	return curConfig,nil
}
