
package main

import "fmt"

// https://leetcode.com/problems/two-sum/description/

// Given an array of integers nums and an integer target, return indices of the
// two numbers such that they add up to target. You may assume that each input
// would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

func twoSum(nums []int, target int) []int {
  // "cache" to remember the numbers we already know have no solutions.
  // Overkill when len(nums) is small, but  alife saver when it's big.
  solutionCache := make(map[int]bool)

  for i := 0 ; i < len(nums) - 1 ; i++ {
    iVal := nums[i]
    // If we remember that the value had no matches, skip it.
    if solutionCache[iVal] {
      continue
    }

    for j := i + 1 ; j < len(nums) ; j++ {
      if iVal + nums[j] == target {
        return []int{ i, j }
      }
    }

    // No match, remember the value.
    solutionCache[iVal] = true
  }

  // As per the problem description, this shoud be an unreachable case.
  return []int{ -1, -1 }
}

func main() {
  // This would take a long ass time without caching the results.
  nums := []int{ 2, 7, 11, 15 }
  target := 9
  fmt.Println(twoSum(nums, target)) // Expected: [ 0, 1 ].

  nums = []int{ 3, 2, 4 }
  target = 6
  fmt.Println(twoSum(nums, target)) // Expected: [ 1, 2 ].

  nums = []int{ 3, 3 }
  target = 6
  fmt.Println(twoSum(nums, target)) // Expected: [ 0, 1 ].

  nums = []int{ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2 }
  target = 4
  fmt.Println(twoSum(nums, target)) // Expected: [ 41, 42 ].
}

