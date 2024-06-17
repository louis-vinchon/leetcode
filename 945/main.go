package main

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/minimum-increment-to-make-array-unique/description/

// You are given an integer array nums. In one move, you can pick an index i
// where 0 <= i < nums.length and increment nums[i] by 1.
//
// Return the minimum number of moves to make every value in nums unique.
//
// The test cases are generated so that the answer fits in a 32-bit integer.
//
// Constraints:
//
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5

// HashMap is too slow.
func minIncrementForUniqueHashMap(nums []int) int {
  counts := make(map[int]int)
  minimum := nums[0]

  for _, value := range nums {
    counts[value]++

    minimum = min(minimum, value)
  }

  // fmt.Printf("counts %v\n", counts)

  numbersDone := 0
  increments := 0

  nextFreeNumber := minimum + 1
  number := minimum
  for numbersDone < len(nums) {
    // fmt.Printf("n: %d\n", n)
    // fmt.Printf("count: %d\n", counts[n])
    if counts[number] == 0 {
      // fmt.Printf("nothing to see here\n---\n")
      number++
      continue
    }
    if counts[number] == 1 {
      // fmt.Printf("just one\n---\n")
      number++
      numbersDone++
      continue
    }

    for {
      nextFreeCandidateCount := counts[nextFreeNumber]

      if nextFreeCandidateCount == 0 {
        break
      }
      nextFreeNumber++
    }

    // fmt.Printf("local max: %v\n---\n", localMax)

    increments += nextFreeNumber - number

    counts[nextFreeNumber]++
    counts[number]--
  }

  return increments
}

func minIncrementForUniqueArray(nums []int) int {
  slices.Sort(nums)

  increments := 0
  for i := 1 ; i < len(nums) ; i++ {
    if nums[i] <= nums[i-1] {
      inc := nums[i-1] - nums[i] + 1
      nums[i] = nums[i-1] + 1

      increments += inc
    }
  }

  return increments
}


func main() {
  function := minIncrementForUniqueArray
  nums := []int{}

  nums = []int{3,2,1,2,1,7}
  fmt.Printf("Result:   %v\n", function(nums))
  fmt.Printf("Expected: %v\n\n", 6)

  nums = []int{1,1,1,2}
  fmt.Printf("Result:   %v\n", function(nums))
  fmt.Printf("Expected: %v\n\n", 5)

  nums = []int{2,0,2}
  fmt.Printf("Result:   %v\n", function(nums))
  fmt.Printf("Expected: %v\n\n", 1)
}



