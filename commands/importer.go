package commands

import (
	"github.com/co0p/neo4ipool/neo4j"
)

type Importer struct {
	Filepath string
	GraphDB  neo4j.GraphDB
}

func (i Importer) Run() (string, error) {

	// parse json

	// create nodes

	// create references

	// push into db

	return "", nil
}
