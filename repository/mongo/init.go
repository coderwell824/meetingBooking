package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"meetingBooking/config"
	"time"
)

var MongoDbClient *mongo.Client

func InitMongoConnection() {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", config.MongoDbAddr, config.MongoDbPort)))
	if err != nil {
		fmt.Println("mongodb连接失败", err)
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("mongodb连接失败", err)
		return
	}
	MongoDbClient = client
	fmt.Println("mongodb连接成功")

}
