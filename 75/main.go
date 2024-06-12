package main

import "fmt"

// https://leetcode.com/problems/sort-colors/description/

// Given an array nums with n objects colored red, white, or blue, sort them
// in-place so that objects of the same color are adjacent, with the colors in
// the order red, white, and blue.
//
// We will use the integers 0, 1, and 2 to represent the color red, white, and
// blue, respectively.
//
// You must solve this problem without using the library's sort function.

// Swap two cells iff needed, slide the window to the right. Repeat.
// This will sort out the bigger numbers first (at the end of the array).
func bubbleSort(nums []int)  {
  for i := 0 ; i < len(nums) ; i++ {
    swapped := false
    for k := 0 ; k < len(nums) - i - 1 ; k++ {
      if nums[k] > nums[k+1] {
        nums[k], nums[k+1] = nums[k+1], nums[k]
        swapped = true
      }
    }

    if !swapped {
      break
    }
  }
}

// Count the occurences of each number. And generate the final array.
// Effective when there are few possible numbers compared to the length of the array.
// It's not the real counting sort because I didn't need stability. So I took some shortcuts.
func countingSort(nums []int)  {
  counts := make([]int, 3)

  for _, val := range nums {
    counts[val]++
  }

  sum := 0
  for i := 0 ; i < 3 ; i++ {
    for k := 0 ; k < counts[i] ; k++ {
      nums[sum + k] = i
    }

    sum += counts[i]
  }
}

var sortColors = bubbleSort

func main() {
  function := countingSort
  nums := []int{}

  nums = []int{2,0,2,1,1,0}
  function(nums)
  fmt.Printf("Res      %v\n", nums)
  fmt.Printf("Expected %v\n", []int{0,0,1,1,2,2})
}

