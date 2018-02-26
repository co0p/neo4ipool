package commands

import (
	"github.com/co0p/neo4ipool/neo4j"
)

type Purger struct {
	GraphDB neo4j.GraphDB
}

func (p *Purger) Run() (string, error) {

	if err := p.GraphDB.Purge(); err != nil {
		return "", err
	}

	return "Successful purged all data", nil
}
