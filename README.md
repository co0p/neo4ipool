neo4ipool
=========

Provide topic detection of news articles based on entities and their relationships. Based on the Nicole White's work found at https://portal.graphgist.org/graph_gists/movie-recommendations-with-k-nearest-neighbors-and-cosine-similarity .

commands
--------

 * ```./neo4ipool -import <path/to/import/json>``` will extract nodes and their relationships and create them in the db
 * ```./neo4ipool -purge``` will remove any data found in the database
 * ```./neo4ipool -topic <path/to/entities>``` will suggest any detected topics based on the entities
 
For examples of the json files please look in the examples folder.
 


installation
------------

You need to have **go** and go's **dep tool** installed. 

    go get -v https://github.com/co0p/neo4ipool
    dep ensure
    go build cmd/neo4ipool/neo4ipool.go
    ./neo4ipool # run the binary ...



current state
-------------

Currently the following node types are extracted from the json (see examples/articles.json for an example). An __Article__ belongsTo a __Category__ and all other NodeTypes are mentionedIn __Article__.
```go
const (
    Article      NodeType = "Article"
    Event        NodeType = "Event"
    Location     NodeType = "Location"
    Person       NodeType = "Person"
    Category     NodeType = "Category"
    Keyword      NodeType = "Keyword"
    Organisation NodeType = "Organisation"
)
```

![graph of nodes](https://raw.githubusercontent.com/co0p/neo4ipool/master/docs/articles_with_tags_and_categories.png)


Fun Cyphers
===========
    
common nodes
-------------

For the articles __article6__ and __article7__, give me all Nodes, that they have in common with the relationship __mentioned_in__.

```sql
MATCH (a1:Article {name:"article6"})<-[r1:mentioned_in]-(n)-[r2:mentioned_in]->(a2:Article {name: "article7"})
RETURN n.name as NAME, r1.weight AS `A1 Weight`, r2.weight AS `A2 Weight`
``` 

Will give you:
```
╒══════════╤═══════════╤═══════════╕
│"NAME"    │"A1 Weight"│"A2 Weight"│
╞══════════╪═══════════╪═══════════╡
│"Südkorea"│"7.7029176"│"11.521997"│
├──────────┼───────────┼───────────┤
│"Berlin"  │"1.265252" │"1.4625446"│
├──────────┼───────────┼───────────┤
│"China"   │"118.79045"│"5.573722" │
└──────────┴───────────┴───────────┘
```

With this information, we can calculate the similarity:

```
                (7.7029176 * 11.521997)  +  (1.265252 * 1.4625446)  +  (118.79045 *5.573722)  
sim(a1,a2) =  ---------------------------------------------------------------------------------------------
                sqrt(7.7029176^2 + 1.265252^2 + 118.79045^2) * sqrt(11.521997^2 + 1.4625446^2 + 5.573722^2)

                752.7084248                              752.7084248
sim(a1,a2) =  ----------------------------------   =   -----------------
                119.0466581   *   12.8826172            1533.6325252


sim(a1,a2) = 0.490801
```

Well, a1 and a2 are not that similar. One could use this approach to add a new relationship between articles such as *similarity*.  


```sql
    MATCH (p1:Article)<-[x:mentioned_in]-(n)-[y:mentioned_in]->(p2:Article)
    WITH SUM(x.weight * y.weight) AS xyDotProduct,
    SQRT(REDUCE(xDot = 0.0, a IN COLLECT(x.weight) | xDot + a^2)) AS xLength,
    SQRT(REDUCE(yDot = 0.0, b IN COLLECT(y.weight) | yDot + b^2)) AS yLength,
    p1, p2
    MERGE (p1)-[s:SIMILARITY]-(p2)
    SET s.similarity = xyDotProduct / (xLength * yLength)
```
