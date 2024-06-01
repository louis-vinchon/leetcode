package main

import "fmt"

// https://leetcode.com/problems/rotate-image/description/

func debug(matrix [][]int) {
  fmt.Println("---")
  for _, row := range matrix {
    fmt.Printf("%v\n", row)
  }
  fmt.Println("---")
}

// The strategy is to rotate one onion layer at a time.
// Tricky.
func rotate(matrix [][]int)  {
  matrixLength := len(matrix)
  layerLength := matrixLength

  for i := 0 ; i < len(matrix) / 2 ; i++ {
    // fmt.Printf("Onion %d\n", i + 1)
    // fmt.Printf("Layer length %d\n", layerLength)
    // We do row length - 1 because the last element of the row / column is
    // also the first element of another column / row, which we already
    // rotated.
    for j:= i ; j < i + layerLength - 1 ; j++ {
      // fmt.Printf("J %d\n", j)
      firstRowPos := &matrix[i][j]
      firstColPos := &matrix[matrixLength - j - 1][i]
      lastRowPos := &matrix[matrixLength - i - 1][matrixLength - j - 1]
      lastColPos := &matrix[j][matrixLength - i - 1]

      temp := *firstRowPos

      // fmt.Printf("firstRowPos value %d\n", *firstRowPos)
      // fmt.Printf("firstColPos %d\n", *firstColPos)
      // fmt.Printf("lastRowPos %d\n", *lastRowPos)
      // fmt.Printf("lastColPos %d\n", *lastColPos)

      *firstRowPos = *firstColPos
      *firstColPos = *lastRowPos
      *lastRowPos = *lastColPos
      *lastColPos = temp
    }

    layerLength -= 2
  }
}

func main() {
  matrix := [][]int { { 1, 2 }, { 3, 4 } }
  debug(matrix)
  rotate(matrix)
  debug(matrix)

  // matrix = [][]int { { 1, 2, 3 }, { 4, 5, 6 }, { 7, 8, 9 } }
  // debug(matrix)
  // rotate(matrix)
  // debug(matrix)

  matrix = [][]int{ { 5, 1, 9, 11 }, { 2, 4, 8, 10 }, { 13, 3, 6, 7 }, { 15, 14, 12, 16 } }
  debug(matrix)
  rotate(matrix)
  debug(matrix)
}

