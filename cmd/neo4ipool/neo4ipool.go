package main

import (
	"flag"
	"log"
	"os"

	"github.com/co0p/neo4ipool/commands"
	"github.com/co0p/neo4ipool/graphdb"
)

var (
	dbURI, importCmd, topicCmd string
	purgeCmd                   bool
)

func init() {
	flag.StringVar(&dbURI, "uri", "http://neo4j:foobar@localhost:7474/db/data", "uri of neo4j location, like https://<user>:<pwd>@host:port")
	flag.StringVar(&importCmd, "import", "", "path to json file for import")
	flag.StringVar(&topicCmd, "topic", "", "path to json file to detect topic for")
	flag.BoolVar(&purgeCmd, "purge", false, "if true, it will delete all data")
}

func main() {
	flag.Parse()

	graphdb, err := graphdb.Connect(dbURI)
	if err != nil {
		log.Fatalf("failed to connect to %s: %s", dbURI, err.Error())
	}

	importer := commands.Import{Filepath: importCmd, GraphDB: &graphdb}
	topicDetector := commands.TopicDetect{Filepath: topicCmd, GraphDB: &graphdb}
	purger := commands.Purge{GraphDB: &graphdb}

	var cmd commands.Runner
	if len(importCmd) > 0 {
		cmd = &importer
	} else if len(topicCmd) > 0 {
		cmd = &topicDetector
	} else if purgeCmd {
		cmd = &purger
	} else {
		flag.Usage()
		os.Exit(0)
	}

	if cmdErr := cmd.Run(); cmdErr != nil {
		log.Fatalf("failed to run command: %s", cmdErr.Error())
	}
}
