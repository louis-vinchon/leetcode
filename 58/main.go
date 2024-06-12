
package main

import "fmt"

// https://leetcode.com/problems/length-of-last-word/

// Given a string s consisting of words and spaces, return the length of the
// last word in the string.
//
// A word is a maximal substring consisting of non-space characters only.

func lengthOfLastWord(s string) int {
  isTrailingSpace := true
  trailingSpaces := 0
  for i := 0 ; i < len(s) ; i++ {
    if s[len(s) - i - 1] == ' ' {
      if !isTrailingSpace {
        return i - trailingSpaces
      } else {
        trailingSpaces++
      }
    } else {
      isTrailingSpace = false
    }
  }

  return len(s) - trailingSpaces
}

func main() {
  function := lengthOfLastWord
  s := ""

  s = "two words"
  fmt.Printf("Result   %v\n", function(s))
  fmt.Printf("Expected %v\n\n", len("words"))

  s = "two"
  fmt.Printf("Result   %v\n", function(s))
  fmt.Printf("Expected %v\n\n", len("two"))

  s = "   fly me   to   the moon  "
  fmt.Printf("Result   %v\n", function(s))
  fmt.Printf("Expected %v\n\n", len("moon"))

  s = "trailing      "
  fmt.Printf("Result   %v\n", function(s))
  fmt.Printf("Expected %v\n\n", len("trailing"))

  s = "t      "
  fmt.Printf("Result   %v\n", function(s))
  fmt.Printf("Expected %v\n\n", len("t"))
}

