package main

import (
	"authentication/cmd"
	"authentication/data"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

func main() {

	log.Println("starting authentication service")

	//Connect to DB
	conn := cmd.ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to postgre")
	}

	//Set up config
	app := cmd.Config{
		DB:     conn,
		Models: data.New(conn),
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
