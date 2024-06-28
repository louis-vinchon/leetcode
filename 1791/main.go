package main

import "fmt"

// https://leetcode.com/problems/find-center-of-star-graph/description/

// There is an undirected star graph consisting of n nodes labeled from 1 to n.
// A star graph is a graph where there is one center node and exactly n - 1
// edges that connect the center node with every other node.
//
// You are given a 2D integer array edges where each edges[i] = [ui, vi]
// indicates that there is an edge between the nodes ui and vi. Return the
// center of the given star graph.

func findCenter(edges [][]int) int {
  if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
    return edges[0][0]
  } else {
    return edges[0][1]
  }
}

func main() {
  function := findCenter
  edges := [][]int{}

  edges = [][]int{
    { 1, 2 },
    { 1, 3 },
  }
  fmt.Printf("Result  : %v\n", function(edges))
  fmt.Printf("Expected: %v\n\n", 1)

  edges = [][]int{
    { 4, 8 },
    { 9, 4 },
  }
  fmt.Printf("Result  : %v\n", function(edges))
  fmt.Printf("Expected: %v\n\n", 4)
}

