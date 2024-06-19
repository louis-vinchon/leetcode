package main

import "fmt"

// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/description/

// You are given an integer array bloomDay, an integer m and an integer k.
//
// You want to make m bouquets. To make a bouquet, you need to use k adjacent
// flowers from the garden.
//
// The garden consists of n flowers, the ith flower will bloom in the
// bloomDay[i] and then can be used in exactly one bouquet.
//
// Return the minimum number of days you need to wait to be able to make m
// bouquets from the garden. If it is impossible to make m bouquets return -1.
//
//
//
// Example 1:
// Input: bloomDay = [1,10,3,10,2], m = 3, k = 1
// Output: 3
// Explanation: Let us see what happened in the first three days. x means flower bloomed and _ means flower did not bloom in the garden.
// We need 3 bouquets each should contain 1 flower.
// After day 1: [x, _, _, _, _]   // we can only make one bouquet.
// After day 2: [x, _, _, _, x]   // we can only make two bouquets.
// After day 3: [x, _, x, _, x]   // we can make 3 bouquets. The answer is 3.
//
// Example 2:
// Input: bloomDay = [1,10,3,10,2], m = 3, k = 2
// Output: -1
// Explanation: We need 3 bouquets each has 2 flowers, that means we need 6 flowers. We only have 5 flowers so it is impossible to get the needed bouquets and we return -1.
//
// Example 3:
// Input: bloomDay = [7,7,7,7,12,7,7], m = 2, k = 3
// Output: 12
// Explanation: We need 2 bouquets each should have 3 flowers.
// Here is the garden after the 7 and 12 days:
// After day 7: [x, x, x, x, _, x, x]
// We can make one bouquet of the first three flowers that bloomed. We cannot make another bouquet from the last three flowers that bloomed because they are not adjacent.
// After day 12: [x, x, x, x, x, x, x]
// It is obvious that we can make two bouquets in different ways.
//
//
//
// Constraints:
//
//     bloomDay.length == n
//     1 <= n <= 10^5
//     1 <= bloomDay[i] <= 10^9
//     1 <= m <= 10^6
//     1 <= k <= n


// Binary search solution (dichotomie en Francais).
// Forget about grouping flowers by k or whatever.
// Just find the biggest and smallest bloom waits. Take the half point. Count
// how many you can make, move the half point until you get the result you want
// (m).
//
// Note to self: you suck at recognizing binary search use cases don't you?
// Don't answer that.

func getNumberOfBouquets(bloomDay []int, k int, day int) int {
  fmt.Printf("get bouquets\n")
  bouquets := 0

  count := 0
  for i := 0 ; i < len(bloomDay) ; i++ {
    if bloomDay[i] <= day {
      count++
    } else {
      count = 0
    }

    if count == k {
      bouquets++
      count = 0
    }
  }

  return bouquets
}

func minDays(bloomDay []int, m int, k int) int {
  maxBloom := bloomDay[0]
  minBloom := bloomDay[0]
  for _, bloom := range bloomDay {
    minBloom = min(minBloom, bloom)
    maxBloom = max(maxBloom, bloom)
  }

  middle := (minBloom + maxBloom) / 2
  // prevMiddle := 0

  minWait := -1
  // for minBloom < maxBloom && prevMiddle != middle {
  for minBloom <= maxBloom {
    // prevMiddle = middle

    // Switch to incremental mode.
    inc := (maxBloom - minBloom) < 6

    fmt.Printf("min %v\n", minBloom)
    fmt.Printf("max %v\n", maxBloom)
    bouquets := getNumberOfBouquets(bloomDay, k, middle)
    fmt.Printf("day %d , inc %t, %d bouquets possible\n", middle, inc, bouquets)

    if bouquets < m {
      print("not enough bouquets, increase\n")
      minBloom = middle

      if inc {
        minBloom += 1
        middle += 1
      } else {
        middle = (middle + maxBloom) / 2
      }
    } else {
      minWait = middle

      print("enough or more than enough, decrease\n")
      maxBloom = middle
      if inc {
        maxBloom -= 1
        middle -= 1
      } else {
        middle = (minBloom + middle) / 2
      }
    }
  }

  return minWait
}

func main() {
  function := minDays
  bloomDays := []int{}
  m := 0
  k := 0

  bloomDays = []int{1,10,2,10,3}
  m = 3
  k = 1
  fmt.Printf("Result  : %v\n", function(bloomDays, m, k))
  fmt.Printf("Expected: %v\n\n", 3)

  bloomDays = []int{1}
  m = 2
  k = 1
  fmt.Printf("Result  : %v\n", function(bloomDays, m, k))
  fmt.Printf("Expected: %v\n\n", -1)

  bloomDays = []int{7,7,7,7,12,7,7}
  m = 2
  k = 3
  fmt.Printf("Result  : %v\n", function(bloomDays, m, k))
  fmt.Printf("Expected: %v\n\n", 12)

  bloomDays = []int{7,7,7,7,12,7,7}
  m = 2
  k = 3
  fmt.Printf("Result  : %v\n", function(bloomDays, m, k))
  fmt.Printf("Expected: %v\n\n", 12)
}


