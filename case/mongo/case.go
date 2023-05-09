package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoClient() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(""))
	if err != nil {
		panic(err)
	}

	return client
}

type EventInfo struct {
	Id           string `bson:"_id,omitempty"`
	UserId       int64  `bson:"userId"`
	Text         string
	Latlng       string
	Country      string
	City         string
	Location     string
	Like         []int64
	LikeCount    map[int64]int64 `bson:"likeCount"`
	ActionUserId int64           `proto:"-" bson:"-"`
	CreateTime   int64           `bson:"createTime"`
	ReplyCount   int64           `bson:"replyCount"`
	UpdateTime   int64           `bson:"updateTime,omitempty"`
	Status       int32           `bson:"status"`
	Remark       string          `bson:"remark,omitempty"`
	Viewers      []int64
}

func test1(db *mongo.Database) {
	c := db.Collection("event")

	match := bson.M{
		"$match": bson.M{
			"status": 1,
		},
	}

	newField := bson.M{
		"$addFields": bson.M{
			"test": bson.M{
				"$cond": bson.M{
					"if": bson.M{
						"$and": []bson.M{
							{"$in": []interface{}{"$userId", []int64{8873}}},
							{"$ne": []interface{}{"$viewers", 1}},
						},
					},
					"then": 1,
					"else": 2,
				},
			},
		},
	}

	sort := bson.M{
		"$sort": bson.M{
			"test": 1,
		},
	}

	// skip := bson.M{}
	limit := bson.M{
		"$limit": 3,
	}

	pipeline := []bson.M{match, newField, sort, limit}

	cur, err := c.Aggregate(context.Background(), pipeline)
	if err != nil {
		fmt.Println(err)
	}

	var res []*EventInfo
	if err = cur.All(context.Background(), &res); err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}

type UserHighlight struct {
	UserId     int64   `bson:"_id,omitempty"`
	Highlights []int64 `bson:"highlights,omitempty"`
}

func test2(db *mongo.Database) {
	c := db.Collection("userHighlight")
	u := &UserHighlight{UserId: 8873, Highlights: []int64{}}
	_, err := c.InsertOne(context.Background(), u)

	res, err := c.UpdateOne(context.Background(),
		bson.M{
			"_id": 8873,
		},
		bson.M{
			"$pull": bson.M{
				"highlights": bson.M{
					"$in": []int64{1, 3},
				},
			},
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	if res.MatchedCount == 0 {
		fmt.Println(err)
	}

}

func test3(db *mongo.Database) {
	c := db.Collection("userHighlight")
	writes := []mongo.WriteModel{
		mongo.NewUpdateOneModel().SetFilter(bson.D{{"_id", 8873}}).SetUpdate(bson.D{{"$push", bson.D{{"follows", 8874}}}}),
		mongo.NewUpdateOneModel().SetFilter(bson.D{{"_id", 8874}}).SetUpdate(bson.D{{"$push", bson.D{{"fans", 8873}}}}),
	}

	c.BulkWrite(context.Background(), writes)

}

func test4(db *mongo.Database) {
	c := db.Collection("userHighlight")
	err := c.FindOne(context.Background(),
		bson.M{
			"_id": 123,
		}).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println(err)
		}
	}
}

func test5(db *mongo.Database) {
	c := db.Collection("club")
	res, err := c.Find(context.Background(),
		bson.M{},
		options.Find())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println(err)
		}
	}

	fmt.Println(res)
}

func Run() {
	client := mongoClient()
	db := client.Database("user_stage")

	test5(db)
}
