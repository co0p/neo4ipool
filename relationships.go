package neo4ipool

type RelationshipType string

const (
	BelongsTo   RelationshipType = "belongs_to"
	MentionedIn RelationshipType = "mentioned_in"
)

type Relationship struct {
	From   Node
	To     Node
	Type   RelationshipType
	Weight float32
}
