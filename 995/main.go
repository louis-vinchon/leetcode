package main

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips/description/

// You are given a binary array nums and an integer k.
//
// A k-bit flip is choosing a subarray of length k from nums and simultaneously
// changing every 0 in the subarray to 1, and every 1 in the subarray to 0.
//
// Return the minimum number of k-bit flips required so that there is no 0 in
// the array. If it is not possible, return -1.
//
// A subarray is a contiguous part of an array.
//
// Constraints:
//
//     1 <= nums.length <= 105
//     1 <= k <= nums.length


// Bruteforce. TLE.
func minKBitFlipsBruteForce(nums []int, k int) int {
  count := 0

  // Just flip when we find a zero.
  for i := 0 ; i <= len(nums) - k ; i++ {
    if nums[i] == 0 {
      count++
      for m := i ; m < i + k ; m++ {
        nums[m] ^= 1
      }
    }
  }

  // Check if we have zeroes left a the end of the array (the only place where
  // they can be).
  if slices.Index(nums[len(nums)-k:], 0) != -1 {
    return -1
  }

  return count
}

// Smarter solution. Don't actually flip the bits. Just remember where the
// flips should happend.
//
// We keep a list of where we flipped bits and then drop them as they become
// irrelevant (out of the k window).
// If there are 2 flips in the window, then it's as if we didn't flip. So if
// the number is 0, it needs a flip, and if it's 1 it doesn't need one.
func minKBitFlipsDeque(nums []int, k int) int {
  flippedBits := []int{}

  count := 0

  for i := 0 ; i < len(nums) ; i++ {
    // Remove old flipped bits as they become irrelevant.
    if len(flippedBits) > 0 && flippedBits[0] == i-k {
      flippedBits = flippedBits[1:]
    }

    // Flip if needed.
    if len(flippedBits) % 2 == nums[i] {
      // Check if we overflow the end of the array.
      if i+k > len(nums) {
        return -1
      }

      // Remember where we flipped.
      flippedBits = append(flippedBits, i);
      count++
    }
  }

  return count
}

func main() {
  function := minKBitFlipsDeque
  nums := []int{}
  k := 0

  nums = []int{1,0,1}
  k = 2
  fmt.Printf("Result  : %v\n", function(nums, k))
  fmt.Printf("Expected: %v\n\n", -1)

  nums = []int{0,0,0}
  k = 3
  fmt.Printf("Result  : %v\n", function(nums, k))
  fmt.Printf("Expected: %v\n\n", 1)

  nums = []int{0,1,0}
  k = 1
  fmt.Printf("Result  : %v\n", function(nums, k))
  fmt.Printf("Expected: %v\n\n", 2)

  nums = []int{0,0,0,1,0,1,1,0}
  k = 3
  fmt.Printf("Result  : %v\n", function(nums, k))
  fmt.Printf("Expected: %v\n\n", 3)

  nums = []int{0,1,0,1,0,1,0,1,0,1,0}
  k = 3
  fmt.Printf("Result  : %v\n", function(nums, k))
  fmt.Printf("Expected: %v\n\n", 6)

  nums = []int{0,1,0,1,0,1,0,1,0,1,0}
  k = 3
  fmt.Printf("Result  : %v\n", function(nums, k))
  fmt.Printf("Expected: %v\n\n", 6)
}

