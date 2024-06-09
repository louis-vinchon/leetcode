package main

import (
	"fmt"
	"math"
)

// HARD
// https://leetcode.com/problems/closest-subsequence-sum/description/

//You are given an integer array nums and an integer goal.
//
// You want to choose a subsequence of nums such that the sum of its elements
// is the closest possible to goal. That is, if the sum of the subsequence's
// elements is sum, then you want to minimize the absolute difference abs(sum -
// goal).
//
// Return the minimum possible value of abs(sum - goal).
//
// Note that a subsequence of an array is an array formed by removing some
// elements (possibly all or none) of the original array.

func minAbsDifference(nums []int, goal int) int {
  diff := goal

  seen := make(map[int]int)
  for _, num := range nums {
    diff := int(math.Abs(float64(num - goal)))

    if diff == 0 {
      return 0
    }
    if seen[diff] != 0 {
      return 0
    }
  }

  return diff
}

func main() {
  function := minAbsDifference
  nums := []int{}
  goal := 0

  nums = []int{ 1,1,1,1,1,1 }
  goal = 1
  fmt.Printf("%v\n", function(nums, goal))
  fmt.Printf("Expected %v\n\n", 1)
}

