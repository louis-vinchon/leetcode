package main

import (
  "fmt"
  "math"
)

var solutionsCache = make(map[int]int)

func getDeletionStepsForNumber(numb int) int {
  value, ok := solutionsCache[numb]
  if ok {
    return value
  }

  // A particular int can be present any number of times.
  // We try to fit as many of the big deletion operation as possible.
  // And then fill with the small deletions.
  // stepsOfTwo := 0
  // stepsOfThree := numb / 3
  // rest := numb % 3
  //
  // The rest can only be 0, 1, or 2.
  // 0 means perfect fit for "3 deletions operations".
  // 1 means that we used one too many "3 deletions operations". So we use one
  // less "3 deletion operations". 1 + 3 = 4, so we will fill with 2 "2
  // deletions operations".
  // 2 means that it's a perfect fit for one additional "2 deletions operation".
  // switch rest {
  // case 0:
  // case 1:
  //   stepsOfThree -= 1
  //   stepsOfTwo = 2
  // default:
  //   stepsOfTwo = 1
  // }

  // The above code can be simplified as:
  // - if it's a multiple of 3, return val / 3
  // - otherwise, the final result (steps of three + step of two) will always be val / 3 + 1 . i.e. Math.ceil(val / 3)
  solution := int(math.Ceil(float64(numb) / 3.0))
  solutionsCache[numb] = solution

  return solution
}

func minOperations(nums []int) int {
  count := make(map[int]int)

  for _, val := range nums {
    count[val] += 1
  }

  steps := 0
  for _, val := range count {
    // We can only delete two or three items at once, only 1 is not possiple.
    if val == 1 {
      return -1
    }

    steps += getDeletionStepsForNumber(val)
  }

  return steps
}

func main() {
  input := []int{ 2, 3, 3, 2, 2, 4, 2, 3, 4 }
  fmt.Printf("%o\n", minOperations(input))

  // input := []int{ 1, 1, 2, 2 }
  // fmt.Printf("%o\n", minOperations(input))
}

