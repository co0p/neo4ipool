package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/jmcvetta/neoism"
)

var dbURL string
var dbPassword string
var dbUser string

func init() {
	flag.StringVar(&dbURL, "url", "", "url location of neo4j db")
	flag.StringVar(&dbUser, "username", "", "username of db")
	flag.StringVar(&dbPassword, "password", "", "password of db")
}

func main() {
	flag.Parse()

	if len(dbURL) == 0 {
		log.Fatalln("usage: go run main.go -url <neo4j url> -user <user> -password <password>")
	}

	fmt.Println("I AM a cli!")

	data := []string{"https://", dbUser, ":", dbPassword, "@", dbURL}
	urlWithCredentials := strings.Join(data, "")
	db, err := neoism.Connect(urlWithCredentials)

	if err != nil {
		log.Fatalf("failed to connect to db: %s", err.Error())
	}

	n, err := db.CreateNode(neoism.Props{"name": "Captain Kirk"})
	if err != nil {
		log.Printf("failed to create node: %s", err.Error())
	}

	name, _ := n.Property("name")
	log.Printf("CREATED NODE with name: %s\n", name)
}
