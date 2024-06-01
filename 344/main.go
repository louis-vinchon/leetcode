package main

import (
	"fmt"
)

// https://leetcode.com/problems/reverse-string/description/

// This is apparenly faster than iterating until len(s) / 2.
func reverseString(s []byte)  {
  i := 0
  j := len(s) - 1
  for i < j {
    s[i], s[j] = s[j], s[i]
    i++
    j--
  }
}

func main() {
  str := []byte{ 'a', 'b', 'c' }
  reverseString(str)
  fmt.Printf("%v\n", str)
}

