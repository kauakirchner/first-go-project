package mongodb

import (
	"context"
	"os"

	"github.com/kauakirchner/first-go-project/src/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	DATABASE_URI     = "DATABASE_URI"
	DB_USER_DATABASE = "DB_USER_DATABASE"
)

func NewMongoDbConnection(ctx context.Context) (*mongo.Database, error) {
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(os.Getenv(DATABASE_URI)),
	)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	logger.Info("Database connected succesfuly", zap.String("journey", "dbConnect"))
	db := client.Database(os.Getenv(DB_USER_DATABASE))
	return db, nil
}
