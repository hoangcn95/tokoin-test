package users

import (
	"context"
	"time"

	"app/model"

	"go.mongodb.org/mongo-driver/bson"
)

// All ..
func (r Repo) All(condition bson.M) (results []model.Users, err error) {
	// TODO: set timeout 15 second
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	collection := r.Session.GetCollectionV2(r.Collection)
	cur, err := collection.Find(ctx, condition)
	if err != nil {
		return
	}
	// Close the cursor once finished
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		temp := model.Users{}
		err := cur.Decode(&temp)
		if err != nil {
			return nil, err
		}

		results = append(results, temp)
	}

	return
}
