package main

import "fmt"
import "math"

// https://leetcode.com/problems/score-of-a-string/description/

func scoreOfString(s string) int {
  sum := 0
  for i := 0 ; i < len(s) - 1 ; i++ {
    diff := int(s[i]) - int(s[i + 1])
    sum += int(math.Abs(float64(diff)))
  }

  return sum
}

func main() {
  fmt.Println(scoreOfString("a"))
  fmt.Println(scoreOfString("hello"))
  fmt.Println(scoreOfString("zaz"))
}

