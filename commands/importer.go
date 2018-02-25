package commands

import (
	"encoding/json"
	"os"

	"github.com/co0p/neo4ipool/neo4j"
)

type Importer struct {
	Filepath string
	GraphDB  neo4j.GraphDB
}

func (i Importer) Run() (string, error) {

	// parse json
	_, err := parseJSON(i.Filepath)
	if err != nil {
		return "", err
	}

	// create nodes

	// create references

	// push into db

	return "", nil
}

type importData struct {
}

func parseJSON(filePath string) (importData, error) {
	in, err := os.Open(filePath)
	if err != nil {
		return importData{}, err
	}

	var parsed importData
	jsonParser := json.NewDecoder(in)
	if err = jsonParser.Decode(&parsed); err != nil {
		return importData{}, err
	}
	return parsed, nil
}
