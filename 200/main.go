package main

import "fmt"

// https://leetcode.com/problems/number-of-islands/description/

// Given an m x n 2D binary grid grid which represents a map of '1's (land) and
// '0's (water), return the number of islands.
//
// An island is surrounded by water and is formed by connecting adjacent lands
// horizontally or vertically. You may assume all four edges of the grid are
// all surrounded by water.

func debugVisitedGrid(grid [][]byte) {
  for _, row := range grid {
    fmt.Printf("%c\n", row)
  }
  fmt.Printf("-----\n")
}

const VISITED = 'X'
const ISLAND = '1'
const WATER = '0'

// Breadth first "visiting of the island".
func visitIsland(grid *[][]byte, rowI int, colI int) int {
  // Skip visited cells.
  if (*grid)[rowI][colI] == VISITED {
    return 0
  }

  // Skip sea cells.
  if (*grid)[rowI][colI] == WATER {
    return 0
  }

  // Mark as visited.
  (*grid)[rowI][colI] = VISITED
  // debugVisitedGrid(*grid)

  // This is an island, check neighboring cells for moar island.
  if 0 < rowI { // Up.
    // fmt.Printf("up\n")
    visitIsland(grid, rowI - 1, colI)
  }
  if rowI < len(*grid) - 1 { // Down.
    // fmt.Printf("down\n")
    visitIsland(grid, rowI + 1, colI)
  }

  if 0 < colI { // Left.
    // fmt.Printf("left\n")
    visitIsland(grid, rowI, colI - 1)
  }
  if colI < len((*grid)[0]) - 1 { // Right.
    // fmt.Printf("right\n")
    visitIsland(grid, rowI, colI + 1)
  }

  return 1
}

func numIslands(grid [][]byte) int {
  count := 0
  for rowI := 0 ; rowI < len(grid) ; rowI++ {
    for colI := 0 ; colI < len(grid[0]) ; colI++ {
      count += visitIsland(&grid, rowI, colI)
      // debugVisitedGrid(grid)
    }
  }

  return count
}

func main() {
  function := numIslands
  grid := [][]byte{}

  grid = [][]byte{
    { '1','1','1','1','1','1' },
    { '1','0','0','0','0','1' },
    { '1','0','1','1','0','1' },
    { '1','0','1','0','0','1' },
    { '1','0','1','1','1','1' },
  }
  fmt.Printf("%v\n", function(grid))
  fmt.Printf("Expected %v\n\n", 1)

  grid = [][]byte{
    { '1','1','1','1','1','1' },
    { '1','1','1','0','0','0' },
    { '1','1','1','1','0','0' },
    { '1','1','0','0','0','1' },
    { '1','0','1','0','0','1' },
  }
  fmt.Printf("%v\n", function(grid))
  fmt.Printf("Expected %v\n\n", 3)

  grid = [][]byte{
    { '1','1','0','0','0' },
    { '1','1','0','0','0' },
    { '0','0','1','0','0' },
    { '0','0','0','1','1' },
    { '0','0','0','1','1' },
  }
  fmt.Printf("%v\n", function(grid))
  fmt.Printf("Expected %v\n\n", 3)

  grid = [][]byte{
    { '1','1'},
    { '1','1'},
  }
  fmt.Printf("%v\n", function(grid))
  fmt.Printf("Expected %v\n\n", 1)

  grid = [][]byte{
    { '0','1'},
    { '1','0'},
  }
  fmt.Printf("%v\n", function(grid))
  fmt.Printf("Expected %v\n\n", 2)

  grid = [][]byte{
    { '0','0'},
    { '0','0'},
  }
  fmt.Printf("%v\n", function(grid))
  fmt.Printf("Expected %v\n\n", 0)

  grid = [][]byte{
    { '0','0'},
    { '0','1'},
  }
  fmt.Printf("%v\n", function(grid))
  fmt.Printf("Expected %v\n\n", 1)

}

