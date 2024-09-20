package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "8080"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {
	//try to connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	log.Println("Connected to RabbitMQ!")
	
	app := Config{
		Rabbit: rabbitConn,
	}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	//&http.Server is a struct that contains the configuration of the server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),// Address to listen on
		Handler: app.routes(),// Handler to invoke, http.DefaultServeMux if nil
	}

	// start the server
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func connect() (*amqp.Connection, error) {
	//connect to rabbitmq
	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection
	
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println("err")
			return nil, err
		}

		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off")
		time.Sleep(backoff)
		continue
	}

	return connection, nil

}