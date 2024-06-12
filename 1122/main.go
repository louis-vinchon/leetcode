package main

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/relative-sort-array/

// Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all
// elements in arr2 are also in arr1.
//
// Sort the elements of arr1 such that the relative ordering of items in arr1
// are the same as in arr2. Elements that do not appear in arr2 should be
// placed at the end of arr1 in ascending order.

func relativeSortArray(arr1 []int, arr2 []int) []int {
  indices := make(map[int]int)
  for _, val := range arr1 {
    indices[val] = slices.Index(arr2, val)
  }

  fmt.Printf("arr1 %v\n", arr1)
  fmt.Printf("arr2 %v\n", arr2)
  fmt.Printf("indices %v\n", indices)

  for i := 0 ; i < len(arr1) - 1 ; i++ {
    swapped := false

    for k := 0 ; k < len(arr1) - i - 1 ; k++ {
      value1 := &arr1[k]
      value2 := &arr1[k+1]
      index1 := indices[*value1]
      index2 := indices[*value2]
      fmt.Printf("value1 %d\n", *value1)
      fmt.Printf("index1 %d\n", index1)
      fmt.Printf("value2 %d\n", *value2)
      fmt.Printf("index2 %d\n", index2)

      // This means the item was not in arr2, and thus should be placed at the end.
      if index1 < 0 || index2 < 0 {
        if index1 == index2 { // Two items that were not in arr2. Sort by their actual numerical ascending order.
          if *value1 > *value2 {
            *value1, *value2 = *value2, *value1
            swapped = true
          }
        } else if index2 != -1 {
          *value1, *value2 = *value2, *value1
          swapped = true
        }
      } else if index1 > index2 { // Regular bubble sort swap.
        *value1, *value2 = *value2, *value1
        swapped = true
      }
      fmt.Printf("%v\n", arr1)
      fmt.Printf("-\n")
    }

    fmt.Printf("swaped %t\n", swapped)
    fmt.Printf("-----\n")

    if !swapped {
      break
    }
  }

  return arr1
}

func main() {
  function := relativeSortArray
  arr1 := []int{}
  arr2 := []int{}

  arr1 = []int{2,3,1,3,2,4,6,7,9,2,19}
  arr2 = []int{2,1,4,3,9,6}
  fmt.Printf("Result   %v\n", function(arr1, arr2))
  fmt.Printf("Expected %v\n\n", []int{2,2,2,1,4,3,3,9,6,7,19})
}

