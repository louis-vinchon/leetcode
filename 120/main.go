package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/triangle/description/

// Given a triangle array, return the minimum path sum from top to bottom.
//
// For each step, you may move to an adjacent number of the row below. More
// formally, if you are on index i on the current row, you may move to either
// index i or index i + 1 on the next row.

func minimumTotal(triangle [][]int) int {
  if len(triangle) == 1 {
    return triangle[0][0]
  }

  // fmt.Printf("%v\n", triangle)
  for i := len(triangle) - 1 ; i > 0 ; i-- {
    row := triangle[i]

    for j := 0 ; j < len(row) - 1 ; j++ {
      parentCell := triangle[i - 1][j]

      triangle[i - 1][j] = int(math.Min(
        float64(parentCell + triangle[i][j]),
        float64(parentCell + triangle[i][j + 1]),
      ))
    }

    // fmt.Printf("%v\n", triangle)
  }

  return triangle[0][0]
}

func main() {
  triangle := [][]int { { 1 }, { 1, 2 } }
  fmt.Printf("res %d\n", minimumTotal(triangle)) // Expected: 2.

  triangle = [][]int { { 1 }, { 1, 2 }, { 4, 1, 1 } }
  fmt.Printf("res %d\n", minimumTotal(triangle)) // Expected 3.

  triangle = [][]int { { -10 } }
  fmt.Printf("res %d\n", minimumTotal(triangle)) // Expected -10.

  triangle = [][]int { { -3 }, { 2, -1 }, { 1, -2, 1 } }
  fmt.Printf("res %d\n", minimumTotal(triangle)) // Expected -6.

  triangle = [][]int { { 2 }, { 3, 4 }, { 6, 5, 7 }, { 4, 1, 8, 3 } }
  fmt.Printf("res %d\n", minimumTotal(triangle)) // Expected 11 ?
}

