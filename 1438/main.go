package main

import (
	"fmt"
	"math"
	"slices"
)

// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/description/

// Given an array of integers nums and an integer limit, return the size of the
// longest non-empty subarray such that the absolute difference between any two
// elements of this subarray is less than or equal to limit.
//
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     0 <= limit <= 10^9

type Multiset []int

func (multiset Multiset) Max() int {
  return multiset[len(multiset)-1];
}

func (multiset Multiset) Min() int {
  return multiset[0];
}

func (multiset Multiset) AbsDiff() int {
  return int(math.Abs(float64(multiset.Max()) - float64(multiset.Min())))
}

func (multiset *Multiset) Push(item int) {
  index, _ := slices.BinarySearch(*multiset, item);
  *multiset = slices.Insert(*multiset, index, item);
}

func (multiset *Multiset) Erase(item int) {
  index := slices.Index(*multiset, item);
  *multiset = slices.Delete(*multiset, index, index+1);
}

func (multiset *Multiset) Init() {
  slices.Sort(*multiset)
}


func longestSubarray(nums []int, limit int) int {
  beginIndex := 0

  multiset := Multiset{}

  longest := 0

  for endIndex := 0 ; endIndex < len(nums) ; endIndex++ {
    fmt.Printf("---\n")
    fmt.Printf("Nums %v\n", nums)
    fmt.Printf("  b Index %d\n", beginIndex)
    fmt.Printf("End Index %d\n", endIndex)
    multiset.Push(nums[endIndex])
    fmt.Printf("Multiset %v\n", multiset)

    diff := multiset.AbsDiff()
    fmt.Printf("Diff %d\n", diff)

    if diff <= limit {
      longest = max(longest, len(multiset))
      fmt.Printf("longest %d\n", longest)
      continue
    }

    for beginIndex < endIndex {
      fmt.Printf("---\n")
      fmt.Printf("Nums %v\n", nums)
      fmt.Printf("Begin Index %d\n", beginIndex)
      fmt.Printf("    e Index %d\n", endIndex)

      multiset.Erase(nums[beginIndex])
      beginIndex++
      fmt.Printf("Multiset %v\n", multiset)

      diff := multiset.AbsDiff()
      fmt.Printf("Diff %d\n", diff)


      if diff <= limit {
        longest = max(longest, len(multiset))
        fmt.Printf("longest %d\n", longest)
        break
      }
    }
  }


  return longest
}

func main() {
  function := longestSubarray
  nums := []int{}
  limit := 0

  nums = []int{8,2,4,7}
  limit = 4
  fmt.Printf("Result  : %v\n", function(nums, limit))
  fmt.Printf("Expected: %v\n\n", 2)

  nums = []int{10,1,2,4,7,2}
  limit = 5
  fmt.Printf("Result  : %v\n", function(nums, limit))
  fmt.Printf("Expected: %v\n\n", 4)

  nums = []int{4,2,2,2,4,4,2,2}
  limit = 0
  fmt.Printf("Result  : %v\n", function(nums, limit))
  fmt.Printf("Expected: %v\n\n", 3)
}

