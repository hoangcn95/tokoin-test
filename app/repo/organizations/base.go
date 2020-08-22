package organizations

import (
	"context"
	"sync"

	"app/common/config"
	"app/repo"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// Repo type
type Repo repo.Mongo

var (
	instance *Repo
	once     sync.Once
)

// New ..
func New() *Repo {
	once.Do(func() {
		instance = &Repo{
			Session:    config.GetConfig().Mongo.Get("tokoin"),
			Collection: "Organizations",
		}

		// Ensure index
		ctx := context.Background()

		// create keys
		keys := make([]mongo.IndexModel, 0)

		// create option
		indexOpts := options.Index()
		indexOpts.SetBackground(true)

		identityKey := mongo.IndexModel{
			Keys:    bsonx.Doc{{Key: "_id", Value: bsonx.Int32(1)}},
			Options: indexOpts,
		}

		keys = append(keys, identityKey)

		collection := instance.Session.GetCollectionV2(instance.Collection)
		collection.Indexes().CreateMany(ctx, keys)
	})

	return instance
}
