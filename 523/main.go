package main

import (
	"fmt"
)

// https://leetcode.com/problems/continuous-subarray-sum/

// Given an integer array nums and an integer k, return true if nums has a good subarray or false otherwise.

// A good subarray is a subarray where:
//
//     its length is at least two, and
//     the sum of the elements of the subarray is a multiple of k.
//
// Note that:
//
//     A subarray is a contiguous part of the array.
//     An integer x is a multiple of k if there exists an integer n such that x = n * k. 0 is always a multiple of k.
//

// The idea:
// - An array of length 0 means sum == 0, and thus 0 % k == 0. An array of
//   length 0 is the beginning of any array right?
// - We are looking for an array of length > 1, where the sum of the elements %
//   k == 0. Notice how the modulo is the same as above.
// - If you do one iteration over the array, building the prefix sum up, and
//   keeping only the modulo. Then wherever our subarray starts, should have the
//   same moulo as the end. This value is basically the modulo of all the garbage
//   before the array, we don't actually care about it, just that it's present
//   somewhere else.
// - So all we have to do is build a map to remember which modulo happened at
//   which index. And when we find a pair, we are golden.
func checkSubarraySum(nums []int, k int) bool {
  if len(nums) < 2 {
    return false
  }

  // -1 is the correct value for the eventality where a prefix sum is 0.
  modulos := map[int]int{ 0: -1 }

  prefixMod := 0
  for i := 0 ; i < len(nums) ; i++ {
    prefixMod = (prefixMod + nums[i]) % k

    fmt.Printf("i %d | %d\n", i, nums[i])
    fmt.Printf("mod %d\n", prefixMod)

    index, ok := modulos[prefixMod]
    if ok {
      if i - index > 1 {
        return true
      }
    } else {
      modulos[prefixMod] = i
    }
  }

  fmt.Printf("%v\n", nums)
  fmt.Printf("%v\n", modulos)

  return false
}

func main() {
  function := checkSubarraySum
  nums := []int{}
  k := 0

  // nums = []int{23,2,4,6,7}
  // k = 6
  // fmt.Printf("%v\n", function(nums, k)) // Expected: true.
  //
  // nums = []int{23,2,6,4,7}
  // k = 6
  // fmt.Printf("%v\n", function(nums, k)) // Expected: true.
  //
  // nums = []int{23,2,6,4,7}
  // k = 13
  // fmt.Printf("%v\n", function(nums, k)) // Expected: false.
  //
  // nums = []int{5,0,0,0}
  // k = 3
  // fmt.Printf("%v\n", function(nums, k)) // Expected: true.

  nums = []int{23,2,4,6,6}
  k = 7
  fmt.Printf("%v\n", function(nums, k)) // Expected: true.
}

