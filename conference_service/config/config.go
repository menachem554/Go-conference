package config

import (
	"time"

	"github.com/spf13/viper"
)

func SetConfig() {
	viper.SetDefault("Port", "9090")
	viper.SetDefault("MongoConnString", "mongodb://localhost:27017/conferenceDB")
	viper.SetDefault("conferenceCollection", "orders")
	viper.SetDefault("MongoClientConnectionTimout", 5*time.Second)
	viper.SetDefault("MongoClientPingTimeout", 5*time.Second)
	viper.SetDefault("MongoCreateIndexTimeout", 5*time.Second)
	viper.SetDefault("GrpcMaxReceivedMessageSize", 16777216)

	viper.BindEnv("Port", "PORT")
	viper.BindEnv("MongoConnString", "MONGO_CONNECTION_STRING")
	viper.BindEnv("conferenceCollection", "MONGO_CONFERENCE_COLLECTION")
	viper.BindEnv("MongoClientConnectionTimout", "MONGO_CLIENT_CONNECTION_TIMEOUT")
	viper.BindEnv("MongoClientPingTimeout", "MONGO_CLIENT_PING_TIMEOUT")
	viper.BindEnv("MongoCreateIndexTimeout", "MONGO_CREATE_INDEX_TIMEOUT")
	viper.BindEnv("GrpcMaxReceivedMessageSize", "GRPC_MAX_RECEIVED_MESSAGE_SIZE")
}