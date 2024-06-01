package main

import (
	"fmt"
	"slices"
)

// 1https://leetcode.com/problems/minimum-falling-path-sum/description/

// Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.
//
// A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right. Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).
//
// func maximumANDSum(nums []int, numSlots int) int {

func minFallingPathSum(matrix [][]int) int {
  for i := len(matrix) - 2 ; i >= 0 ; i-- {
    for j := 0 ; j < len(matrix) ; j++ {
      minimum := matrix[i][j]

      // Check the extremities.
      if j > 0 {
        minimum = min(minimum, matrix[i + 1][j - 1])
      }
      if j < len(matrix) - 1 {
        minimum = min(minimum, matrix[i + 1][j + 1])
      }

      matrix[i][j] = minimum
    }
  }

  return slices.Min(matrix[0])
}

func main() {
  nums := [][]int{ { 2, 1, 3 }, { 6, 5, 4 }, { 7, 8, 9 } }
  fmt.Printf("%v\n\n", minFallingPathSum(nums)) // Expected: 13.

  nums = [][]int{ { -19, 57 }, { -40, -5 } }
  fmt.Printf("%v\n\n", minFallingPathSum(nums)) // Expected: -59.
}

