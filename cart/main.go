package main

import (
	"cart/common"
	"cart/controller"
	"cart/domain/repository"
	"cart/domain/service"
	"cart/proto/cart"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracingv2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

var QPS = 100

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "ferry")
	if err != nil {
		log.Error(err)
	}

	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		//注册中心地址
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//链路追踪
	trace,io,err := common.NewTracer("go.micro.service.cart","localhost:6831")
	if err!=nil {
		fmt.Println(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(trace)

	//服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		//设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		//添加consul注册中心
		micro.Registry(consulRegistry),
		//绑定链路追踪
		micro.WrapHandler(opentracingv2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)), // var QPS = 100
	)
	//初始化服务
	srv.Init()


	//获取mysql配置,func中是配置中心的k-v的映射路径中不用带前缀
	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig,"mysql")

	//初始化数据库 "root:123456@/micro?charset=utf8&parseTime=True&loc=Local"
	db,err := gorm.Open("mysql",mysqlConfig.User+":"+mysqlConfig.Pwd+"@/"+
		mysqlConfig.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//禁止复表
	db.SingularTable(true)
	//初始化操作数据库层面的repository
	cartRepository := repository.NewCartRepository(db)
	//执行一次，初始化创建数据库表
	/*cartRepository.InitTable()*/



	//初始化service
	cartService := service.NewCartDataService(cartRepository)
	//服务绑定指定Service
	err = cart.RegisterCartHandler(srv.Server(), &controller.CartController{CartService: cartService})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
