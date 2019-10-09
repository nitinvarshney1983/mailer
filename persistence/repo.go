package persistence

import mongo "go.mongodb.org/mongo-driver/mongo"

//Repo is the type that contains collection refrence
type Repo struct {
	col *mongo.Collection
}
