package inits

import (
	"context"
	"fmt"
	"github.com/yuanzhichao-star/public_pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InitMongoDB() {
	//设置用户选项
	data := config.AppCong.MongoDB
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d/?compressors=snappy,zlib,zstd",
		data.User, data.Password, data.Host, data.Port))
	//mongodb连接
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	//检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	log.Println("mongodb init success")
	//设置存储文档库
	config.Coll = client.Database("mongodb").Collection("user_auth")
}
