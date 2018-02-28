package commands

import (
	"encoding/json"
	"log"
	"os"

	"github.com/co0p/neo4ipool"
	"github.com/co0p/neo4ipool/graphdb"
)

type entry struct {
	ArticleID   string      `json:"id"`
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

	articles := []neo4ipool.Node{}
	categories := []neo4ipool.Node{}
	belongsTos := []neo4ipool.Relationship{}

	// create nodes
	for _, v := range parsedEntries {
		article := newNode(neo4ipool.Article, v.ArticleID)
		category := newNode(neo4ipool.Category, v.Category)
		belongsTo := newRelationship(neo4ipool.BelongsTo, article, category)

		articles = append(articles, article)
		categories = append(categories, category)
		belongsTos = append(belongsTos, belongsTo)
	}

	if err := i.GraphDB.CreateNodes(neo4ipool.Article, articles); err != nil {
		return err
	}
	log.Printf("created %d articles", len(articles))

	if err := i.GraphDB.CreateNodes(neo4ipool.Category, categories); err != nil {
		return err
	}
	log.Printf("created %d categories", len(categories))

	if err := i.GraphDB.CreateRelationships(neo4ipool.BelongsTo, belongsTos); err != nil {
		return err
	}
	log.Printf("created %d belongTo relationships", len(belongsTos))
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

func newNode(t neo4ipool.NodeType, n string) neo4ipool.Node {
	return neo4ipool.Node{Type: t, Name: n}
}

func newRelationship(t neo4ipool.RelationshipType, from neo4ipool.Node, to neo4ipool.Node) neo4ipool.Relationship {
	return neo4ipool.Relationship{Type: t, From: from, To: to}
}
