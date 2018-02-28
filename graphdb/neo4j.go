package graphdb

import (
	"github.com/co0p/neo4ipool"
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

func (c *GraphDB) CreateNode(n neo4ipool.Node) error {
	return nil
}

func (c *GraphDB) CreateRelationship(r neo4ipool.Relationship) error {
	return nil
}

func (c *GraphDB) Purge() error {
	cq := neoism.CypherQuery{
		Statement: `MATCH (n)
		DETACH DELETE n
		RETURN n.name
		`,
		Parameters: nil,
		Result:     nil,
	}

	return c.client.Cypher(&cq)
}
