package commands

import (
	"encoding/json"
	"log"
	"os"

	"github.com/co0p/neo4ipool/graphdb"
)

type entry struct {
	Category    string      `json:"category"`
	Linguistics linguistics `json:"linguistics"`
}

type linguistics struct {
	Events   []linguisticsEntry `json:"events"`
	Geos     []linguisticsEntry `json:"geos"`
	Keywords []linguisticsEntry `json:"keywords"`
	Orgs     []linguisticsEntry `json:"orgs"`
	Persons  []linguisticsEntry `json:"persons"`
}

type linguisticsEntry struct {
	Lemma  string   `json:"lemma"`
	Token  []string `json:"token"`
	Weight float32  `json:"weight"`
}

type Import struct {
	Filepath string
	GraphDB  *graphdb.GraphDB
}

func (i *Import) Run() error {

	// parse json
	parsedEntries, err := i.parseJSON()
	if err != nil {
		return err
	}

	log.Printf("found %d entries in '%s'", len(parsedEntries), i.Filepath)

	// create nodes

	// create references

	// push into db

	return nil
}

func (i Import) parseJSON() ([]entry, error) {
	in, err := os.Open(i.Filepath)
	if err != nil {
		return nil, err
	}

	var parsed []entry
	jsonParser := json.NewDecoder(in)
	if err = jsonParser.Decode(&parsed); err != nil {
		return nil, err
	}
	return parsed, nil
}
