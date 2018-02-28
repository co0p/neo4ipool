neo4ipool
=========

Provide topic detection of news articles based on entities and their relationships.

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

Currently the nodes __article__ and __cateogory__ are supported. They are connected with the relationship __belongs_to__. Parsing a json with 100 articles and creating the graph currently takes about 1 second on my 1.6 Ghz machine.

![graph of nodes](https://raw.githubusercontent.com/co0p/neo4ipool/master/docs/articles_belong_to_categories.png)
