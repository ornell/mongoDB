package mongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	eh "github.com/ornell/errorhandling"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Content interface {

}

func MongoConnect(constring, db, collect string)(mgocol *mongo.Collection){
	clientOptions := options.Client().ApplyURI(constring)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	eh.ErrorPrint(err, "Can't connect to mongoDB")

	err = client.Ping(context.TODO(), nil)

	eh.ErrorPrint(err, "Can't connect to mongoDB")

	collection := client.Database(db).Collection(collect)
	return collection
}

func MongoInsert(collection *mongo.Collection, document interface{})(){
	insertResult, err := collection.InsertOne(context.TODO(), document)
	eh.ErrorPanic(err, "Can't connect to mongoDB")
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
}

func MongoFindOne(collection *mongo.Collection, filter, result interface{})interface{}{
	fmt.Println(filter)
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	eh.ErrorPanic(err, "Can't connect to mongoDB")
	return result
}

func MongoUpdate(collection *mongo.Collection, filter, update interface{})(){
	// MongoUpdate requires filter (same as findone) and a update document 	update := bson.D{		{"$set", bson.D{			{"out", time.Now().Unix()},		}},	}

}

func MongoReturnAll(collection *mongo.Collection, filter bson.M) []*interface{}{
	var Results []*interface{}

	cur, err := collection.Find(context.TODO(), filter)
	eh.ErrorPanic(err, "Can't connect to mongoDB")
	for cur.Next(context.TODO()){
		var content interface{}
		err = cur.Decode(&content)
		eh.ErrorPrint(err, "Can't connect to mongoDB")
		Results = append(Results, &content)
	}
	return Results
}