package main

import (
  "fmt"
)

// https://leetcode.com/problems/set-mismatch/description/

// You have a set of integers s, which originally contains all the numbers from
// 1 to n. Unfortunately, due to some error, one of the numbers in s got
// duplicated to another number in the set, which results in repetition of one
// number and loss of another number.
//
// You are given an integer array nums representing the data status of this set
// after the error.
//
// Find the number that occurs twice and the number that is missing and return
// them in the form of an array.

// nums[i] > 1
func findErrorNums(nums []int) []int {
  // Array for counting the elements we have seen. We don't need a hashmap
  // because we know how many elements we expect and the data is dense (it
  // spans from 1 to n, with exactly one exception)
  //
  // + 1 to the size because we are not going to use the first "cell" for the
  // first element (index = 0 for n = 1). If we did, we'd have to do a - 1 for
  // every single element of the array. So that would be n times more
  // expensive thatn doing this simple + 1.
  counts := make([]int, len(nums) + 1)

  // We can find the duplicate sa we go.
  duplicatedNum := 0
  for _, num := range nums {
    counts[num]++

    if duplicatedNum == 0 && counts[num] > 1 {
      duplicatedNum = num
    }
  }

  // fmt.Printf("%v\n", counts)

  // We find the missing element. Do not forget to start at index 1!
  omittedNum := 0
  for i := 1 ; i < len(counts) ; i++ {
    if counts[i] == 0 {
      omittedNum = i
      break
    }
  }

  return []int{ duplicatedNum, omittedNum }
}


func main() {
  nums := []int{ 1, 2, 2, 4 }
  fmt.Printf("%v\n", findErrorNums(nums)) // Expected [2, 3].

  // input := []int{ 1, 1, 2, 2 }
  // fmt.Printf("%o\n", minOperations(input))
}

