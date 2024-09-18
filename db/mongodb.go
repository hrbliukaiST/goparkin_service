go
package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "myproject/config"
)

var MongoDB *mongo.Client

func InitMongoDB() {
    var err error
    MongoDB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.AppConfig.MongoDBURI))
    if err != nil {
        log.Fatal(err)
    }
    if err := MongoDB.Ping(context.TODO(), nil); err != nil {
        log.Fatal(err)
    }
}