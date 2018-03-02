package neo4ipool

type NodeType string

const (
	Article      NodeType = "Article"
	Event        NodeType = "Event"
	Location     NodeType = "Location"
	Person       NodeType = "Person"
	Category     NodeType = "Category"
	Keyword      NodeType = "Keyword"
	Organisation NodeType = "Organisation"
)

type Node struct {
	Type NodeType
	Name string
}
