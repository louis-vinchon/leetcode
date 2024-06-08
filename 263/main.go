package main

import (
  "fmt"
)

// https://leetcode.com/problems/ugly-number/

// An ugly number is a positive integer which does not have a prime factor other than 2, 3, and 5.
//
// Given an integer n, return true if n is an ugly number.


// It's faster when they are in this particular order.
var divs = []int{ 5, 3, 2 }

var memo = make(map[int]bool)

func isUgly(n int) bool {
  if n <= 0 {
    return false
  }
  if n == 1 {
    return true
  }

  val, ok := memo[n]
  if ok {
    return val
  }

  isUg := false
  for _, div := range divs {
    rest := n % div

    if rest == 0 {
      isUg = isUg || isUgly(n / div)
    }
  }

  memo[n] = isUg

  return isUg
}

func main() {
  function := isUgly
  input := 15
  fmt.Printf("%d | %t\n", input, function(input)) // Expected true.

  input = 13
  fmt.Printf("%d | %t\n", input, function(input)) // Expected false.

  input = -2
  fmt.Printf("%d | %t\n", input, function(input)) // Expected false.

  input = 6
  fmt.Printf("%d | %t\n", input, function(input)) // Expected true.
}

