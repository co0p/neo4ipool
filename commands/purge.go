package commands

import (
	"github.com/co0p/neo4ipool/graphdb"
)

type Purge struct {
	GraphDB *graphdb.GraphDB
}

func (p *Purge) Run() error {
	return p.GraphDB.Purge()
}
