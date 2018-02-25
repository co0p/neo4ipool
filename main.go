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

	graphdb, err := neo4j.Connect(dbURI)
	if err != nil {
		log.Fatalf("failed to connect to %s: %s", dbURI, err.Error())
	}

	importer := commands.Importer{Filepath: importCmd, GraphDB: graphdb}
	topicDetector := commands.TopicDetector{Filepath: topicCmd, GraphDB: graphdb}
	purger := commands.Purger{GraphDB: graphdb}

	var res string
	var cmdErr = errors.New("")

	if len(importCmd) > 0 {
		res, cmdErr = importer.Run()
	} else if len(topicCmd) > 0 {
		res, cmdErr = topicDetector.Run()
	} else if purgeCmd {
		res, cmdErr = purger.Run()
	} else {
		flag.Usage()
		os.Exit(0)
	}

	if cmdErr != nil {
		log.Fatalf("failed to run command: %s", cmdErr.Error())
	}
	log.Println(res)
}
