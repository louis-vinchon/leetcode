package main

import (
  "fmt"
  "slices"
)

// https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/

// You are given a positive integer n representing the number of nodes of a
// Directed Acyclic Graph (DAG). The nodes are numbered from 0 to n - 1
// (inclusive).
//
// You are also given a 2D integer array edges, where edges[i] = [fromi, toi]
// denotes that there is a unidirectional edge from fromi to toi in the graph.
//
// Return a list answer, where answer[i] is the list of ancestors of the ith
// node, sorted in ascending order.
//
// A node u is an ancestor of another node v if u can reach v via a set of
// edges.
//
// Constraints:
//     1 <= n <= 1000
//     0 <= edges.length <= min(2000, n * (n - 1) / 2)
//     edges[i].length == 2
//     0 <= fromi, toi <= n - 1
//     fromi != toi
//     There are no duplicate edges.
//     The graph is directed and acyclic.

func processAncestors(directAncestors *[][]int, index int, processed *[]bool) {
  if (*processed)[index] {
    return
  }

  allAncestors := (*directAncestors)[index]
  for _, ancestor := range (*directAncestors)[index] {
    processAncestors(directAncestors, ancestor, processed)

    allAncestors = slices.Concat(allAncestors, (*directAncestors)[ancestor])
  }

  slices.Sort(allAncestors)
  allAncestors = slices.Compact(allAncestors)
  (*directAncestors)[index] = allAncestors
  (*processed)[index] = true
}

func getAncestors(n int, edges [][]int) [][]int {
  directAncestors := make([][]int, n)

  // fmt.Printf("directAncestors %v\n", directAncestors)

  for _, edge := range edges {
    directAncestors[edge[1]] = append(directAncestors[edge[1]], edge[0])
  }

  processed := make([]bool, n+1)
  for i := range directAncestors {
    if processed[i] {
      continue
    }

    processAncestors(&directAncestors, i, &processed)
  }

  // fmt.Printf("directAncestors %v\n", directAncestors)
  return directAncestors
}

func main() {
  function := getAncestors
  n := 0
  edges := [][]int{}

  // n = 4
  // edges = [][]int{{0,1}, {1,2}, {2,3}}
  // fmt.Printf("Result  : %v\n", function(n, edges))
  // fmt.Printf("Expected: %v\n\n", [][]int{{}, {0}, {0, 1}, {0, 1, 2}})

  n = 8
  edges = [][]int{{0,3},{0,4},{1,3},{2,4},{2,7},{3,5},{3,6},{3,7},{4,6}}
  fmt.Printf("Result  : %v\n", function(n, edges))
  fmt.Printf("Expected: %v\n\n", [][]int{{},{},{},{0,1},{0,2},{0,1,3},{0,1,2,3,4},{0,1,2,3}})
}

