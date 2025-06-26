package mongo

import (
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// DOCS CON COPILOT U OTRO? para todas las funciones

// pasar MONGO_DB_URI para que sea testeable?

func getConnection() (*mongo.Client, error) {
	dbURI := os.Getenv("MONGO_DB_URI")
	if dbURI == "" {
		return nil, errors.New("environment variable MONGO_DB_URI is not set")
	}

	timeDuration := time.Second * 10

	cp := options.ClientOptions{
		ConnectTimeout: &timeDuration,
	}

	c, err := mongo.Connect(&cp, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}
	return c, nil
	// se puede?, ver microblog
	// mongo.Connect(&cp, options.Client().ApplyURI(dbURI))
}
