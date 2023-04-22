package main

import (
	"context"
	"fmt"
	"log"
	"logger/cmd"
	"logger/data"
	"net/http"
	"net/rpc"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	webPort  = "80"
	gRpcPort = "50001"
)

var client *mongo.Client

func main() {
	//connect to Mongo
	mongoClient, err := cmd.ConnctToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	//Create a context inorder to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := cmd.Config{
		Models: data.New(client),
	}

	//Register the rpc server
	err = rpc.Register(new(cmd.RPCServer))

	go app.RpcListen()

	//Start the webserver
	log.Println("Starting service at port: ", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
