package app

import (
	eventbus2 "code.byted.org/ad/open_core/queue/eventbus"
	mysql "code.byted.org/ad/open_core/storage/db/mysql/v2"
	"code.byted.org/ad/open_core/storage/es"
	"code.byted.org/ad/open_core/storage/redis"
	eventbus "code.byted.org/eventbus/client-go"
	"code.byted.org/gopkg/env"
	"code.byted.org/gopkg/lang/strings"
	"gorm.io/gorm"
)

var (
	RedisCli      *redis.Client
	CreativeDBCli *gorm.DB
	EsCli         *es.Client
	//DynamicConfigGetter = &DynamicConfig{}
	NotificationSendProducer eventbus.Producer
)

func InitClients() {
	//initRedis()
	//// initEs()
	initCreativeDB()
	initProducer()
	//initRPC()
	//initTCC()
}

//	func initRedis() {
//		var err error
//		RedisCli, err = redis.NewClient(&ServiceConf.Redis)
//		if err != nil {
//			panic(err)
//		}
//	}
func initCreativeDB() {
	var err error
	dbConfig := &ServiceConf.CreativeDB

	CreativeDBCli, err = mysql.NewClient(dbConfig)
	if err != nil {
		panic(err)
	}
}

func initProducer() {
	var err error
	if !strings.IsBlank(ServiceConf.NotificationSendProducer.EventName) && !env.IsBoe() {
		NotificationSendProducer, err = eventbus2.NewProducer(&ServiceConf.NotificationSendProducer)
		if err != nil {
			panic(err)
		}
	}
}

func CloseProducer() {
	if NotificationSendProducer != nil {
		_ = NotificationSendProducer.Close()
	}
}

