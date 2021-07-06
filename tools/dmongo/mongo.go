package dmongo

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// mongodb Client
var Client *mongo.Client

func InitDB() {
	var err error
	uri := "mongodb://" + viper.GetString("database.mongodb.username") + ":" + viper.GetString("database.mongodb.password") + "@" + viper.GetString("database.mongodb.host") + ":" + viper.GetString("database.mongodb.port") + "/" + viper.GetString("database.mongodb.database")
	// all connect or select/query timeout
	timeout := viper.GetDuration("database.mongodb.timeout") * time.Second
	clientOptions := options.Client().ApplyURI(uri).SetServerSelectionTimeout(timeout)
	Client, err = mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// return Default config mongodb database
func DefaultDB() *mongo.Database {
	return Client.Database(viper.GetString("database.mongodb.database"))
}
