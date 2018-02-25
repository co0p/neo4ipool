package neo4j

import (
	"github.com/jmcvetta/neoism"
)

type GraphDB struct {
	client *neoism.Database
}

func Connect(uri string) (GraphDB, error) {

	db := GraphDB{}
	client, err := neoism.Connect(uri)
	db.client = client
	return db, err
}
