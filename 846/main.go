package main

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/hand-of-straights/

// Alice has some number of cards and she wants to rearrange the cards into
// groups so that each group is of size groupSize, and consists of groupSize
// consecutive cards.
//
// Given an integer array hand where hand[i] is the value written on the ith
// card and an integer groupSize, return true if she can rearrange the cards,
// or false otherwise.

// 1. Count cards.
// 1. Make a list of all card types. Sort them. Infer the minimum and maximum card values.
// 1. Iterate over the cards. Pick the smallest one, try to complete a group
//    from there. Decrease the counts as you go. update the minimum if we
//    exhausted a card type. Don't forget to check the maximum.
func isNStraightHand(hand []int, groupSize int) bool {
  fmt.Printf("hand %v\n", hand)
  if len(hand) % groupSize != 0 {
    return false
  }

  // minimum := hand[0]
  // maximum := minimum
  counts := make(map[int]int)
  for _, val := range hand {
    counts[val]++

    // minimum = min(minimum, val)
    // maximum = max(maximum, val)
  }

  keys := []int{}
  for key := range counts {
    keys = append(keys, key)
  }
  slices.SortFunc(keys, func(i, j int) int {
    return i - j
  })

  minI := 0
  minimum := keys[0]
  maximum := keys[len(keys)-1]


  // fmt.Printf("counts %v\n", counts)
  // fmt.Printf("min %d\n", minimum)
  // fmt.Printf("max %d\n", maximum)

  grouped := 0
  for grouped != len(hand) {
    // fmt.Printf("grouped so far %d\n", grouped)
    // fmt.Printf("counts %v\n", counts)
    // fmt.Printf("min %d\n", minimum)
    // fmt.Printf("max %d\n", maximum)
    if minimum + groupSize > maximum + 1 {
      // fmt.Printf("not possible min max\n")
      return false
    }

    localMin := minimum
    for i := 0 ; i < groupSize ; i++ {
      // fmt.Printf("dealing card %d\n", localMin + i)
      counts[localMin + i]--
      remaining := counts[localMin + i]
      // fmt.Printf("remaining %d\n", remaining)

      if remaining == 0 {
        minI += 1
        if minI < len(keys) {
          minimum = keys[minI]
        }
      } else if remaining < 0 {
        return false
      }
    }

    grouped += groupSize
  }

  return true
}

func main() {
  hand := []int {1,2,3,6,2,3,4,7,8}
  groupSize := 3
  fmt.Printf("res %t\n", isNStraightHand(hand, groupSize)) // Expected: true.

  hand = []int {1,2,3,4,5}
  groupSize = 4
  fmt.Printf("res %t\n", isNStraightHand(hand, groupSize)) // Expected: false.

  hand = []int {1,2,4}
  groupSize = 3
  fmt.Printf("res %t\n", isNStraightHand(hand, groupSize)) // Expected: false.

  hand = []int {66,75,4,37,92,87,68,65,58,100,97,42,19,66,73,1,5,44,30,29,76,31,12,35,26,93,9,36,90,16,86,15,4,9,13,98,10,14,18,90,89,3,10,65,24,31,43,25,54,55,54,81,10,80,31,12,15,14,59,27,64,93,32,26,69,79,13,32,29,24,27,91,92,82,37,101,100,61,74,30,91,62,36,92,28,23,4,63,55,3,11,11,101,22,34,25,2,75,43,72}
  groupSize = 2
  fmt.Printf("res %t\n", isNStraightHand(hand, groupSize)) // Expected: true.

  hand = []int {1,1,2,2,3,3}
  groupSize = 2
  fmt.Printf("res %t\n", isNStraightHand(hand, groupSize)) // Expected: false.

  hand = []int {1,1,2,2,3,3}
  groupSize = 3
  fmt.Printf("res %t\n", isNStraightHand(hand, groupSize)) // Expected: true.
}

