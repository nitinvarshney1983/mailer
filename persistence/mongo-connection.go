package persistence

import (
	"context"
	"strconv"

	log "github.com/nitinvarshney1983/mailer/logging"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoConnection for mongo connection
//this is a wrapper of monngo session got from mongo driver
type MongoConnection struct {
	client *mongo.Client
}

var instance *MongoConnection

//SetUp mongoConnection
func SetUp(configurations interface{}) {
	cnf, ok := configurations.(map[string]interface{})
	if !ok {
		log.Fatal("no db connection details present")
	}
	log.Info(cnf)
	log.Info(cnf["clientoptions"])
	if cnf["clientoptions"] != nil {
		conf, ok := cnf["clientoptions"].(map[string]interface{})
		if !ok {
			log.Fatal("no mongo db connection details present")
		}
		log.Info(conf)
		port := conf["port"].(int64)
		serverURL := "mongodb://" + conf["server"].(string) + ":" + strconv.FormatInt(port, 10) + "/" + conf["database"].(string)
		log.Info(serverURL)
		clientOptions := options.Client().ApplyURI(serverURL)
		// credentials := options.Credential{
		// 	AuthMechanism: "SCRAM-SHA-1",
		// 	//AuthMechanismProperties: SERVICE_NAME,
		// 	AuthSource:  conf["database"].(string),
		// 	Username:    conf["username"].(string),
		// 	Password:    conf["password"].(string),
		// 	PasswordSet: true,
		// }
		// clientOptions.SetAuth(credentials)
		log.Info(clientOptions)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
		}

		instance.client = client
	}

}

//GetConnection For DB
func GetConnection(dbName string) *mongo.Database {
	return instance.client.Database(dbName)
}
