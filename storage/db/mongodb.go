package db

import (
	"context"
	"fmt"

	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/config"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/logger"
	"gitlab.com/samandarobidovfrd/voxe_transcoding_service/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	MongoConn *mongo.Database
}

func NewConn(cfg *config.Config, log logger.Logger) *DB {
	var (
		connectionTimeout = 3
		mongoConn         *mongo.Client
		err               error
	)

	log.Info("main: MONGOCONFIG",
		logger.String("Host", cfg.MongoDBHost),
		logger.Int("Port", cfg.MongoDBPort),
		logger.String("Database", cfg.MongoDBDatabase),
	)

	credential := options.Credential{
		AuthSource:    cfg.MongoDBDatabase,
		Username:      cfg.MongoDBUser,
		Password:      cfg.MongoDBPassword,
		AuthMechanism: "SCRAM-SHA-256",
	}

	mongoString := fmt.Sprintf("mongodb://%s:%d", cfg.MongoDBHost, cfg.MongoDBPort)

	ctx, cancel := utils.ContextTimoutNSecond(connectionTimeout)
	defer cancel()

	if cfg.Environment == "develop" {
		mongoConn, err = mongo.Connect(
			ctx,
			options.Client().ApplyURI(mongoString),
		)
	} else {
		mongoConn, err = mongo.Connect(
			ctx,
			options.Client().ApplyURI(mongoString).SetAuth(credential),
		)
	}

	if err != nil {
		log.Error("failed to connecto to mongodb", logger.Error(err))
		panic(err)
	}

	if err := mongoConn.Ping(context.TODO(), nil); err != nil {
		log.Error("Cannot connect to database error ->", logger.Error(err))
		panic(err)
	}

	connDB := mongoConn.Database(cfg.MongoDBDatabase)
	log.Info("Connected to MongoDB", logger.Any("database: ", connDB.Name()))

	return &DB{
		MongoConn: connDB,
	}
}
