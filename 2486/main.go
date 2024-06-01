package main

import (
  "fmt"
)

// https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/description/
// Return the minimum number of characters that need to be appended to the end
// of s so that t becomes a subsequence of s.
//
// A subsequence is a string that can be derived from another string by
// deleting some or no characters without changing the order of the remaining
// characters.

func appendCharacters(s string, t string) int {
  tLength := len(t)
  matchingChars := 0
  for i := 0 ; i < len(s) && matchingChars < tLength ; i++ {
    if s[i] == t[matchingChars] {
      matchingChars++
    }
  }

  return tLength - matchingChars
}

func main() {
  s, t := "coaching", "coding"
  fmt.Printf("%v\n", appendCharacters(s, t)) // Should be 4.

  s, t = "abcde", "a"
  fmt.Printf("%v\n", appendCharacters(s, t)) // Should be 0.

  s, t = "z", "abcde"
  fmt.Printf("%v\n", appendCharacters(s, t)) // Should be 5.
}

