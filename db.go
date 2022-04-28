package main

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Email string
	Data  struct {
		Collections map[string][]struct {
			Name     string
			FeedLink string
		}
	}
}

func GetDBConnection() (*mongo.Database, error) {
	ctx, release := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	release()
	if err != nil {
		return nil, err
	}
	db := mongoClient.Database(os.Getenv("SHIPYARD_DATABASE"))

	return db, nil
}

func GetLastFetched(db *mongo.Database) (time.Time, error) {
	ctx, release := context.WithTimeout(context.Background(), 10*time.Second)
	rawMetadata := db.Collection("metadata").FindOne(ctx, bson.M{})
	release()

	var metadata struct {
		LastFetched time.Time
	}
	err := rawMetadata.Decode(&metadata)
	if err != nil {
		return time.Now(), err
	}

	return metadata.LastFetched, nil
}

func SetLastFetched(db *mongo.Database, date time.Time) error {
	ctx, release := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := db.Collection("metadata").UpdateOne(ctx, bson.M{}, bson.M{
		"$set": bson.M{
			"LastFetched": date,
		},
	})
	release()

	return err
}

func GetUsers(db *mongo.Database) ([]User, error) {
	ctx, release := context.WithTimeout(context.Background(), 10*time.Second)
	usersCursor, err := db.Collection("users").Find(ctx, bson.M{})
	release()
	if err != nil {
		panic(err)
	}

	// Load all users
	var users []User
	ctx, release = context.WithTimeout(context.Background(), 60*time.Second)
	err = usersCursor.All(ctx, &users)
	release()
	if err != nil {
		return nil, err
	}

	return users, nil
}
