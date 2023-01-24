# Graph-Library

This program was developed for the subject Graph Theory at Universidade Federeal do Rio de Janeiro (UFRJ). It's purpose is to explore graphs making searches and running some alghorithms.

### Versions
This project was separated in 3 versions, v3 is the current one on the repository and the previous ones can be found at "releases", they have the following differences:
#### v1
Deals with graphs unweighted and undirected

BFS (Breadth-first search)
DFS (Depth-first search)
Find components (The size of each one and the vertexes in it)

#### v2
Deals with weighted directed graphs

Dijkstra: Find the shortest path between 2 vertexes
Prim: Find the Minimum spanning tree in the graph (A tree conecting the whole graph with the minimum sum of weights of its edges)

#### v3
Also deals with weighted directed graphs

Ford Fulkerson: Finds the maximum flow between 2 given vertexes


## 💻 Prerequisite

- A Go compiler
This code does not use any additional library and it's all written in Go


## ☕ Using the program

### Creating the graph to be used

The graph to be analyzed shoud be in a txt file (the filename is up to you) in the following format:

```
4
1 2 2.3
1 4 1.1
3 2 3
2 4 987
```

The number in the first line is the number of vertexes (int).

Each line represent an edge (sourceVertex goalVertex weight) int int float

PS: In v1 the graph is not weighted or directed, so the third value can be disregarded and a vertex from 1 to 2 is the same as a vertex from 2 to 1.

### Running the program

Once your graph is created you are now ready to run the program.

The program will ask you the name of your graph file. And than the methods you want to use, the vertex from and to which you would like it to run...


## Don't wanna flex but...
For the versions 2 and 3 of this project the professor had the students present their work and vote for which they found the best. In both, my project won with all students voting for my project.
