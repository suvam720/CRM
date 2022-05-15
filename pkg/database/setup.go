package database

import (
	"context"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "Users"
const colName = "UserDetails"

var connectionString = func() string {
	myEnv, _ := godotenv.Read()
	dbstr := myEnv["MONGO_URI"]
	return dbstr

}
var Collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString())

	client, _ := mongo.Connect(context.TODO(), clientOption)

	Collection = client.Database(dbName).Collection(colName)
}
