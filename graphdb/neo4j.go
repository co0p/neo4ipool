package graphdb

import (
	"fmt"

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

func (c *GraphDB) CreateNodes(t neo4ipool.NodeType, n []neo4ipool.Node) error {

	type prop map[string]string
	props := []prop{}
	for _, v := range n {

		p := prop{
			"name": v.Name,
			"type": string(v.Type),
		}
		props = append(props, p)
	}

	stmt := `UNWIND $props AS article
	MERGE (:%s { name: article.name, type: article.type })
	`
	cq := neoism.CypherQuery{
		Statement:  fmt.Sprintf(stmt, t),
		Parameters: neoism.Props{"props": props},
		Result:     nil,
	}

	return c.client.Cypher(&cq)
}

func (c *GraphDB) CreateRelationships(t neo4ipool.RelationshipType, r []neo4ipool.Relationship) error {
	type prop map[string]string
	props := []prop{}
	for _, v := range r {

		p := prop{
			"from":      v.From.Name,
			"from_type": string(v.From.Type),
			"to":        v.To.Name,
			"to_type":   string(v.To.Type),
			"weight":    fmt.Sprintf("%v", v.Weight),
		}
		props = append(props, p)
	}

	stmt := `UNWIND $props AS relationship
	MATCH (n{name: relationship.from, type: relationship.from_type}), (m{name: relationship.to, type: relationship.to_type})
	MERGE (n)-[:%s {weight: relationship.weight}]->(m)
	`
	cq := neoism.CypherQuery{
		Statement:  fmt.Sprintf(stmt, t),
		Parameters: neoism.Props{"props": props},
		Result:     nil,
	}

	return c.client.Cypher(&cq)
}

func (c *GraphDB) Purge() error {
	cq := neoism.CypherQuery{
		Statement: `MATCH (n)
		DETACH DELETE n
		`,
		Parameters: nil,
		Result:     nil,
	}

	return c.client.Cypher(&cq)
}
