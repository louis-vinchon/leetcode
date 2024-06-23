package main

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/magnetic-force-between-two-balls/

// In the universe Earth C-137, Rick discovered a special form of magnetic
// force between two balls if they are put in his new invented basket. Rick has
// n empty baskets, the ith basket is at position[i], Morty has m balls and
// needs to distribute the balls into the baskets such that the minimum
// magnetic force between any two balls is maximum.
//
// Rick stated that magnetic force between two different balls at positions x
// and y is |x - y|.
//
// Given the integer array position and the integer m. Return the required
// force.
//
// Constraints:
//     n == position.length
//     2 <= n <= 10^5
//     1 <= position[i] <= 10^9
//     All integers in position are distinct.
//     2 <= m <= position.length

func canFitBalls(position []int, m int, targetMinForce int) bool {
  // Always put a ball in the first (last?) spot as we need to maximise spacing anyway.
  ballsInPlace := 1
  fmt.Printf("%d has a ball\n", position[0])

  rangeBeginning := position[0]

  for i := 1 ; i < len(position) && ballsInPlace < m ; i++ {
    if position[i] - rangeBeginning >= targetMinForce {
      fmt.Printf("%d has a ball\n", position[i])
      ballsInPlace++;
      rangeBeginning = position[i]
    }
  }

  return ballsInPlace == m
}

// _                _ _ _ _ _ _

func maxDistance(position []int, m int) int {
  slices.Sort(position)
  fmt.Printf("position: %v  - balls: %v\n", position, m)

  // Minimum and maximum forces.
  minimum := 1 // Balls next to each other (worse case scenario).
  maximum := (position[len(position)-1] - position[0]) / (m-1) // Balls evenly distributed best case scenario).
  fmt.Printf("max force: %d\n", maximum)

  targetForce := (maximum + minimum) / 2

  lastFit := 0
  for minimum <= maximum {
    fmt.Printf("---\ntarget force: %d\n", targetForce)
    fits := canFitBalls(position, m, targetForce)

    fmt.Printf("fits: %t\n", fits)

    inc := maximum - minimum < 6
    if inc {
      if fits {
        lastFit = targetForce
        minimum++
        targetForce++
      } else {
        maximum--
        targetForce--
      }
    } else {
      if fits {
        lastFit = targetForce
        minimum = targetForce
        targetForce = (targetForce + maximum) / 2
      } else {
        maximum = targetForce
        targetForce = (targetForce + minimum) / 2
      }
    }
  }

  return lastFit
}

func main() {
  function := maxDistance
  position := []int{}
  m := 0

  position = []int{1,2,3,4,7}
  m = 3
  fmt.Printf("Result  : %v\n", function(position, m))
  fmt.Printf("Expected: %v\n\n", 3)

  position = []int{5,4,3,2,1,1000000000}
  m = 2
  fmt.Printf("Result  : %v\n", function(position, m))
  fmt.Printf("Expected: %v\n\n", 999999999)

  position = []int{1,2,3,4,5,10}
  m = 2
  fmt.Printf("Result  : %v\n", function(position, m))
  fmt.Printf("Expected: %v\n\n", 9)

  position = []int{1,2,10}
  m = 3
  fmt.Printf("Result  : %v\n", function(position, m))
  fmt.Printf("Expected: %v\n\n", 1)

  position = []int{1,6,7,8,9,10}
  m = 3
  fmt.Printf("Result  : %v\n", function(position, m))
  fmt.Printf("Expected: %v\n\n", 4)

  position = []int{6,7,8,9,10}
  m = 3
  fmt.Printf("Result  : %v\n", function(position, m))
  fmt.Printf("Expected: %v\n\n", 2)
}

