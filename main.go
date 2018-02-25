package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/co0p/neo4ipool/commands"
	"github.com/co0p/neo4ipool/neo4j"
)

var (
	dbURI, importCmd, topicCmd string
	purgeCmd                   bool
)

func init() {
	flag.StringVar(&dbURI, "uri", "http://neo:neo@localhost:7474/db/data", "uri of neo4j location, like https://<user>:<pwd>@host:port")
	flag.StringVar(&importCmd, "import", "", "path to json file for import")
	flag.StringVar(&topicCmd, "topic", "", "path to json file to detect topic for")
	flag.BoolVar(&purgeCmd, "purge", false, "if true, it will delete all data")
}

func main() {
	flag.Parse()

	_, err := neo4j.Connect(dbURI)
	if err != nil {
		log.Fatalf("failed to connect to %s: %s", dbURI, err.Error())
	}

	var str string
	var cmd string
	var cmdErr = errors.New("")

	if len(importCmd) > 0 {
		cmd = "import"
		str, cmdErr = commands.Import(importCmd)
	} else if len(topicCmd) > 0 {
		cmd = "topic"
		str, cmdErr = commands.DetectTopic(topicCmd)
	} else if purgeCmd {
		cmd = "purge"
		str, cmdErr = commands.Purge()
	} else {
		flag.Usage()
		os.Exit(0)
	}

	if cmdErr != nil {
		log.Fatalf("failed to run %s command: %s", cmd, cmdErr.Error())
	}
	log.Println(str)
}
