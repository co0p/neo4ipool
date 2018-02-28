package neo4ipool

type RelationshipType string

const (
	BelongsTo RelationshipType = "belongs_to"
)

type Relationship struct {
	From Node
	To   Node
	Type RelationshipType
}
