package main

import (
  "fmt"
)

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

func strStr(haystack string, needle string) int {
  res := -1
  needleLength := len(needle)
  for i := 0 ; i <= len(haystack) - needleLength ; i++ {
    if haystack[i:i + needleLength] == needle {
      return i
    }
  }

  return res
}

func main() {
  haystack, needle := "sadbutsad", "sad"
  fmt.Printf("%o\n", strStr(haystack, needle))

  haystack, needle = "leetcode", "code"
  fmt.Printf("%o\n", strStr(haystack, needle))

  haystack, needle = "leetcode", "leeto"
  fmt.Printf("%o\n", strStr(haystack, needle))
}

