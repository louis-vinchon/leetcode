package main

import "fmt"

// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/

// Given a valid parentheses string s, return the nesting depth of s. The
// nesting depth is the maximum number of nested parentheses.

// Example 1:
// Input: s = "(1+(2*3)+((8)/4))+1"
// Output: 3
// Explanation: Digit 8 is inside of 3 nested parentheses in the string.

// Not much to be said about this one.
func maxDepth(s string) int {
  maxDepth := 0
  currentDepth := 0

  for _, char := range s {
    if char == '(' {
      currentDepth++
      if currentDepth > maxDepth {
        maxDepth = currentDepth
      }
    } else if char == ')' {
      currentDepth--
    }
  }

  return maxDepth
}

func main() {
  input := "(1+(2*3)+((8)/4))+1"
  fmt.Println(fmt.Sprintf("%d", maxDepth(input))) // Expected 3.
}


