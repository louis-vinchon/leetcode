package main

import (
  "fmt"
  "math"
)

// https://leetcode.com/problems/perfect-squares/

// Given an integer n, return the least number of perfect square numbers that
// sum to n.
//
// A perfect square is an integer that is the square of an integer; in other
// words, it is the product of some integer with itself. For example, 1, 4, 9,
// and 16 are perfect squares while 3 and 11 are not.

// Preseed with a couple well known solutions to save time.
var memo = map[int]int{0:0, 1:1, 4:1}

func solve(squares []int, n int) int {
  val, ok := memo[n]
  if ok {
    return val
  }

  solution := math.MaxInt
  // Start with the biggest squares, because we are greedy.
  for i := len(squares) - 1 ; i >= 0 ; i-- {
    if squares[i] > n {
      continue
    } else if n == squares[i] {
      solution = 1
      break
    } else {
      solution = min(solution, 1 + solve(squares, n - squares[i]))
    }
  }

  memo[n] = solution
  // fmt.Printf("memo %v\n", memo)

  return solution
}

func numSquares(n int) int {
  if n == 1 {
    return 1
  }

  squares := []int{}
  for i := 1 ; ; i++ {
    square := i*i
    if square > n {
      break
    }
    squares = append(squares, square)
  }

  // fmt.Printf("squares %v\n", squares)

  return solve(squares, n)
}

func main() {
  input := 12
  fmt.Printf("%v\n", numSquares(input)) // Expected: 3. (4 + 4 + 4).

  input = 13
  fmt.Printf("%v\n", numSquares(input)) // Expected: 2. (9 + 4).

  input = 22
  fmt.Printf("%v\n", numSquares(input)) // Expected: 3. (9 + 9 + 4).

  input = 121
  fmt.Printf("%v\n", numSquares(input)) // Expected: 1. (121 aka 11 * 11).

  input = 625
  fmt.Printf("%v\n", numSquares(input)) // Expected: 1. (625 aka 25 * 25).

  input = 500
  fmt.Printf("%v\n", numSquares(input)) // Expected: 2. (484 + 16).
}

