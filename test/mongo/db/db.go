package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	DB struct {
		*mongo.Database
	}
)
