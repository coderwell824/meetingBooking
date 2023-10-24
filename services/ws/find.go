package ws

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"meetingBooking/repository/mongo"
	"meetingBooking/repository/mongo/model"
	"sort"
	"time"
)

type SendSortMessage struct {
	Content   string `json:"content"`
	Read      uint   `json:"read"`
	CreatedAt int64  `json:"created_at"`
}

func insertMessage(database, id, content string, read uint, expireTime int64) error {
	//将消息插入到mongoDB中
	collection := mongo.MongoDbClient.Database(database).Collection(id)
	comment := model.Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expireTime,
		Read:      read,
	}

	_, err := collection.InsertOne(context.Background(), comment)
	return err
}

func FindMany(databaseName, sendId, id string, time int64, pageSize int) (results []model.Result, err error) {

	var resultId []model.Trainer
	var resultSendId []model.Trainer

	idCollection := mongo.MongoDbClient.Database(databaseName).Collection(id)
	sendIdCollection := mongo.MongoDbClient.Database(databaseName).Collection(sendId)
	//排序
	sendIdTimeCurrentSort, err := sendIdCollection.Find(context.TODO(),
		options.Find().SetSort(bson.D{{"startTime", -1}}),
		options.Find().SetLimit(int64(pageSize)))

	idTimeCurrentSort, err := idCollection.Find(context.TODO(),
		options.Find().SetSort(bson.D{{"startTime", -1}}),
		options.Find().SetLimit(int64(pageSize)))

	err = sendIdTimeCurrentSort.All(context.TODO(), &resultSendId)
	err = idTimeCurrentSort.All(context.TODO(), &resultId)
	results, _ = AppendAndSort(resultId, resultSendId)
	return
}

// FirstFindMsg 首次查询(将对方发来的所有未读消息全部读取出来)
func FirstFindMsg(databaseName, sendId, id string) (results []model.Result, err error) {
	var resultMe []model.Trainer
	var resultYou []model.Trainer

	sendIdCollection := mongo.MongoDbClient.Database(databaseName).Collection(sendId)
	idCollection := mongo.MongoDbClient.Database(databaseName).Collection(sendId)

	filter := bson.M{"read": bson.M{
		"&all": []uint{0},
	}}

	sendIdCurrentSort, err := sendIdCollection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{
		"startTime", 1}}), options.Find().SetLimit(1))

	if sendIdCurrentSort == nil {
		return
	}

	var unReadMessages []model.Trainer
	err = sendIdCurrentSort.All(context.TODO(), &unReadMessages)
	if err != nil {
		log.Println("sendIdCurrentSort error", err)
	}

	if len(unReadMessages) > 0 {
		timeFilter := bson.M{"startTime": bson.M{
			"$gte": unReadMessages[0].StartTime,
		}}
		sendIdTimeCurrentSort, _ := sendIdCollection.Find(context.TODO(), timeFilter)
		idTimeCurrentSort, _ := sendIdCollection.Find(context.TODO(), timeFilter)

		err = sendIdTimeCurrentSort.All(context.TODO(), &resultYou)
		err = idTimeCurrentSort.All(context.TODO(), &resultMe)
		results, err = AppendAndSort(resultMe, resultYou)
	} else {
		results, err = FindMany(databaseName, sendId, id, 999999999, 10)
	}
	overTimeFilter := bson.D{
		{"$and", bson.A{
			bson.D{{"endTime", bson.M{"&lt": time.Now().Unix()}}},
			bson.D{{"read", bson.M{"$eq": 1}}},
		}},
	}
	_, _ = sendIdCollection.DeleteMany(context.TODO(), overTimeFilter)
	_, _ = idCollection.DeleteMany(context.TODO(), overTimeFilter)
	// 将所有的维度设置为已读
	_, _ = sendIdCollection.UpdateMany(context.TODO(), filter, bson.M{
		"$set": bson.M{"read": 1},
	})
	_, _ = sendIdCollection.UpdateMany(context.TODO(), filter, bson.M{
		"&set": bson.M{"ebdTime": time.Now().Unix() + int64(3*expirationTime)},
	})

	return
}

func AppendAndSort(resultId, resultSendId []model.Trainer) (results []model.Result, err error) {
	for _, r := range resultId {
		sendSortMsg := SendSortMessage{ //构造返回的msg
			Content:   r.Content,
			Read:      r.Read,
			CreatedAt: r.StartTime,
		}

		result := model.Result{ //构造返回所有的内容
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSortMsg),
			From:      "me",
		}

		results = append(results, result)
	}
	for _, r := range resultSendId {
		sendSortMsg := SendSortMessage{ //构造返回的msg
			Content:   r.Content,
			Read:      r.Read,
			CreatedAt: r.StartTime,
		}

		result := model.Result{ //构造返回所有的内容
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSortMsg),
			From:      "you",
		}

		results = append(results, result)
	}
	sort.Slice(results, func(i, j int) bool { return results[i].StartTime < results[j].StartTime })
	return
}
