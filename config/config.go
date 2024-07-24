package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/ini.v1"
)

var (
	MongoClient *mongo.Client
	Http        string

	MongoName     string
	MongoAddress  string
	MongoPassword string
	MongoPort     string

	RedisAddress  string
	RedisPassword string
	RedisName     string
)

func init() {
	file, err := ini.Load("./config/cfg.ini")
	if err != nil {
		panic("Error loading config: " + err.Error())
	}

	//load
	loadDbInfo(file)
	loadRedisInfo(file)
	loadServerInfo(file)
}

func loadDbInfo(file *ini.File) {
	MongoName = file.Section("Mongo").Key("MongoName").String()
	MongoAddress = file.Section("Mongo").Key("MongoAddress").String()
	MongoPassword = file.Section("Mongo").Key("MongoPassword").String()
	MongoPort = file.Section("Mongo").Key("MongoPort").String()
}

func loadRedisInfo(file *ini.File) {
	RedisAddress = file.Section("Redis").Key("RedisAddress").String()
	RedisPassword = file.Section("Redis").Key("RedisPassword").String()
}

func loadServerInfo(file *ini.File) {
	Http = file.Section("service").Key("Http").String()

}
