# Graph-Project
Graph theory college project

# Graph-Library

This program was developed for the subject Graph Theory at Universidade Federeal do Rio de Janeiro (UFRJ). It's purpose is to explore graphs making searches and running some alghorithms.

### Versions
This project was separated in 3 versions, v3 is the current one and the previous ones can be found at "releases", they have the following differences:
#### v1
Deals with graphs unweighted and undirected

BFS (Breadth-first search
DFS (Depth-first search)
Find components (The size of each one and the vertexes on it)

#### v2
Deals with weighted directed graphs

Dijkstra: Find the cheapest path between 2 vertexes)
Prim: Find the Minimum spanning tree in the graph (A tree conecting the whole graph with the minimum possible sum of the weights of its edges)

#### v3
Also deals with weighted directed graphs

Ford Fulkerson: Finds the maximum flow between 2 given vertexes


## 💻 Prerequisite

- A Go compiler
This code does not use any additional library and it's all written in Go


## ☕ Executing the program

The graph to be analyzed shoud be in a txt file in the following format:

```
4
1 2 2.3
1 4 1.1
3 2 3
2 4 999
```

The number in the first is the number of vertexes
Each line represent an edge (sourceVertex goalVertex weight)

PS: In v1 the graph is not weighted or directed, so the third column can be disconsidered and a vertex from 1 to 2 is the same as a vertex from 2 to 1



