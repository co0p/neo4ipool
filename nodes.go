package neo4ipool

type NodeType string

const (
	Article  NodeType = "Article"
	Event    NodeType = "Event"
	Location NodeType = "Location"
	Person   NodeType = "Person"
	Category NodeType = "Category"
)

type Node struct {
	Type NodeType
	Name string
}
