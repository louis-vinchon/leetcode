package main

import "fmt"

// https://leetcode.com/problems/patching-array/description/

// Given a sorted integer array nums and an integer n, add/patch elements to
// the array such that any number in the range [1, n] inclusive can be formed
// by the sum of some elements in the array.
//
// Return the minimum number of patches required.
//
// Constraints:
//
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^4
//     nums is sorted in ascending order.
//     1 <= n <= 2^31 - 1

// This one is all about noticing that, if you can make (sum) all numbers until S,
// then if the next number X in the list is inferior or equal to S + 1, you an
// then that means you can make all numbers until S + X.
//
// e.g. If you have {1,2,3}, you can make all number until 6.
//      If you have {1,2,3,7}, you can make all number until 6 + 7 == 13.
//      Because {1,2,3} make 1-6, 7 does 7, and all numbers 7-13 can me made
//      with 7 + whatever combination of {1,2,3}.
//      7 is 6 + 1.
//
// If however we had {1,2,3,8}. Then 7 cannot be made. We could only make 1-6, 8-14.
func minPatches(nums []int, n int) int {
  patches := 0

  maxPossibleSum := 0
  numsIterator := 0

  // As long as we didn't manage to make all numbers until n.
  for maxPossibleSum < n {
    if numsIterator < len(nums) && nums[numsIterator] <= maxPossibleSum+1 {
      maxPossibleSum += nums[numsIterator]
      numsIterator++
    } else {
      patches++
      maxPossibleSum += maxPossibleSum + 1
    }
  }

  return patches
}

func main() {
  function := minPatches
  nums := []int{}
  n := 0

  nums = []int{1,3}
  n = 6
  fmt.Printf("Result  : %v\n", function(nums, n))
  fmt.Printf("Expected: %v\n\n", 1)

  nums = []int{1,2,10,20}
  n = 46
  fmt.Printf("Result  : %v\n", function(nums, n))
  fmt.Printf("Expected: %v\n\n", 2)

  nums = []int{1,5}
  n = 600
  fmt.Printf("Result  : %v\n", function(nums, n))
  fmt.Printf("Expected: %v\n\n", 8)
}

