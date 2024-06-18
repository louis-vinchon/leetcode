package main

import "fmt"

// https://leetcode.com/problems/sum-of-square-numbers/description/

// Given a non-negative integer c, decide whether there're two integers a and b
// such that a^2 + b^2 = c.
//
// Constraints:
//     0 <= c <= 2^31 - 1


// Far from optimal but it works for leetcode and it's easy to follow.
// I'll call it a day.
func judgeSquareSum(c int) bool {
  memo := make(map[int]bool)

  lastSquare := 0
  for i := 0 ; lastSquare <= c ; i++ {
    // fmt.Printf("%d\n", i)
    square := i*i
    diff := c - square

    if memo[diff] || square == diff {
      return true
    }

    memo[square] = true
    lastSquare = square

    // fmt.Printf("%v\n", memo)
  }

  return false
}

func main() {
  function := judgeSquareSum
  c := 0

  c = 5
  fmt.Printf("Result  : %v\n", function(c))
  fmt.Printf("Expected: %v\n\n", true)

  c = 3
  fmt.Printf("Result  : %v\n", function(c))
  fmt.Printf("Expected: %v\n\n", false)

  c = 146 // 11*11 + 5*5
  fmt.Printf("Result  : %v\n", function(c))
  fmt.Printf("Expected: %v\n\n", true)

  c = 8 // 4*4 + 4*4
  fmt.Printf("Result  : %v\n", function(c))
  fmt.Printf("Expected: %v\n\n", true)

  c = 128476148
  fmt.Printf("Result  : %v\n", function(c))
  fmt.Printf("Expected: %v\n\n", true)

  c = 4
  fmt.Printf("Result  : %v\n", function(c))
  fmt.Printf("Expected: %v\n\n", true)

  c = 0
  fmt.Printf("Result  : %v\n", function(c))
  fmt.Printf("Expected: %v\n\n", true)
}

