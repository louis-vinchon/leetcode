package main

import "fmt"

// https://leetcode.com/problems/count-number-of-nice-subarrays/description/

// Given an array of integers nums and an integer k. A continuous subarray is
// called nice if there are k odd numbers on it.
//
// Return the number of nice sub-arrays.
//
// Constraints:
//
//     1 <= nums.length <= 50000
//     1 <= nums[i] <= 10^5
//     1 <= k <= nums.length

func numberOfSubarrays(nums []int, k int) int {
  niceSubArrays := 0
  currentSum := 0
  // Keys == prefixSum, value == occurences.
  prefixSum := map[int]int{ 0: 1 }

  for i := 0; i < len(nums); i++ {
    currentSum += nums[i] % 2;

    // Find subarrays with sum k ending at i.
    sumMinusK, ok := prefixSum[currentSum - k]
    if ok {
      niceSubArrays += sumMinusK
    }

    // Increment the current prefix sum in hashmap.
    prefixSum[currentSum]++;
  }

  return niceSubArrays
}

func main() {
  function := numberOfSubarrays
  nums := []int{}
  k := 0

  nums = []int{1,1,2,1,1}
  k = 3
  fmt.Printf("Result  : %v\n", function(nums, k))
  fmt.Printf("Expected: %v\n\n", 2)

  nums = []int{1}
  k = 1
  fmt.Printf("Result  : %v\n", function(nums, k))
  fmt.Printf("Expected: %v\n\n", 1)

  nums = []int{1,1}
  k = 2
  fmt.Printf("Result  : %v\n", function(nums, k))
  fmt.Printf("Expected: %v\n\n", 1)

  nums = []int{2,4,6}
  k = 1
  fmt.Printf("Result  : %v\n", function(nums, k))
  fmt.Printf("Expected: %v\n\n", 0)

  nums = []int{2,2,2,1,2,2,1,2,2,2}
  k = 2
  fmt.Printf("Result  : %v\n", function(nums, k))
  fmt.Printf("Expected: %v\n\n", 16)
}

