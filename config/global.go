package config

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	AppCong AppConfig              //动态配置
	DB      *gorm.DB               //mysql数据库
	Rdb     *redis.Client          //redis缓存
	Ctx     = context.Background() //上下文联调（redis、mongodb）
	Coll    *mongo.Collection      //mongodb数据库木
	Es      *elasticsearch.Client  //es客户端
)
