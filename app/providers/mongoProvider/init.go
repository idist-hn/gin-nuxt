package mongoProvider

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"idist-core/app/providers/configProvider"
	"idist-core/app/providers/loggerProvider"
	"time"
)

var mongoDB *mongo.Database

var CTimeOut = 60 * time.Second

func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), CTimeOut)
	defer cancel()
	loggerProvider.GetLogger().Info("------------------------------------------------------------")
	loggerProvider.GetLogger().Info("Mongo: Đang khởi tạo kết nối...")
	CTimeOut = time.Duration(configProvider.GetConfig().GetInt64("mongo.timeout")) * time.Second
	c := configProvider.GetConfig()

	var uri = c.GetString("mongo.uri")
	if len(uri) == 0 {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s",
			c.GetString("mongo.username"),
			c.GetString("mongo.password"),
			c.GetString("mongo.host"),
			c.GetString("mongo.port"),
			c.GetString("mongo.database"),
			c.GetString("mongo.authSource"))
	}
	if c.GetString("mongo.username") == "" {
		uri = fmt.Sprintf("mongodb://%s:%s/%s?authSource=%s",
			c.GetString("mongo.host"),
			c.GetString("mongo.port"),
			c.GetString("mongo.database"),
			c.GetString("mongo.authSource"),
		)
	}

	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			if configProvider.GetConfig().GetBool("log.query.detail") == true {
				loggerProvider.GetLogger().Info(evt.Command.String())
			} else {
				switch evt.CommandName {
				case "find":
					loggerProvider.GetLogger().Info("db.getCollection(" + evt.Command.Lookup(evt.CommandName).String() + ").find(" + evt.Command.Lookup("filter").String() + ")")
					return
				case "insert":
					loggerProvider.GetLogger().Info("db.getCollection(" + evt.Command.Lookup(evt.CommandName).String() + ").insert(...)")
					return
				case "update":
					loggerProvider.GetLogger().Info("db.getCollection(" + evt.Command.Lookup(evt.CommandName).String() + ").update(...)")
					return
				default:
					loggerProvider.GetLogger().Info(evt.Command.String())
				}
			}
		},
	}
	clientOptions := options.Client().ApplyURI(uri).SetMonitor(cmdMonitor)
	if client, err := mongo.Connect(ctx, clientOptions); err != nil {
		loggerProvider.GetLogger().Info("Mongo: Lỗi kết nối")
		panic(err)
	} else {
		loggerProvider.GetLogger().Info("Mongo: Khởi tạo kết nối thành công")
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			loggerProvider.GetLogger().Info("Mongo: Ping không thành công", zap.Error(err), zap.String("Mongo uri:", uri))
			panic(err)
		} else {
			loggerProvider.GetLogger().Info("Mongo: {ping: Pong}")
		}
		mongoDB = client.Database(c.GetString("mongo.database"))
	}
}

func GetMongoDB() *mongo.Database {
	return mongoDB
}

func CloseMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), CTimeOut)
	defer cancel()
	_ = mongoDB.Client().Disconnect(ctx)
}
