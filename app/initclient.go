package app

import (
	"gorm.io/gorm"
)

var (
	RedisCli            *redis.Client
	CreativeDBCli       *gorm.DB
	EsCli               *es.Client
	DynamicConfigGetter = &DynamicConfig{}
)

func InitClients() {
	initRedis()
	initEs()
	initCreativeDB()
	initProducer()
	initRPC()
	initTCC()
}

func initRedis() {
	var err error
	RedisCli, err = redis.NewClient(&ServiceConf.Redis)
	if err != nil {
		panic(err)
	}
}

func initDB() {
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
