package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"session26/internal/models"
)

type Mongodb struct {
	db *mongo.Client
}

func (m *Mongodb) CreateUser(ctx context.Context, user models.User) (interface{}, error) {
	coll := m.db.Database("users").Collection("user")
	log.Println("User here ----", user)
	result, err := coll.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return result.InsertedID, nil
}

func (m *Mongodb) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	coll := m.db.Database("users").Collection("user")
	// Creates a query filter to match documents in which the "username" is

	filter := bson.D{{"username", userName}}
	// Retrieves the first matching document
	var result models.User
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m *Mongodb) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	coll := m.db.Database("users").Collection("user")

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *Mongodb) UpdateUser(ctx context.Context, user models.User) error {
	coll := m.db.Database("users").Collection("user")
	filter := bson.D{{"username", user.Username}}
	update := bson.M{"$set": bson.M{"password": user.Password}}
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (m *Mongodb) DeleteUser(ctx context.Context, user models.User) error {
	coll := m.db.Database("users").Collection("user")
	log.Println("User here ----", user)
	//hex, err := primitive.ObjectIDFromHex(user.ID)
	//if err != nil {
	//	return err
	//}
	filter := bson.M{"id": user.ID}
	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

func NewMongodb(db *mongo.Client) DbRepository {
	return &Mongodb{db}
}
