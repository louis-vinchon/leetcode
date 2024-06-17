package main

import (
	"fmt"
	"math"
	"slices"
)

// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/description/

// There are n seats and n students in a room. You are given an array seats of
// length n, where seats[i] is the position of the ith seat. You are also given
// the array students of length n, where students[j] is the position of the jth
// student.
//
// You may perform the following move any number of times:
//
//     Increase or decrease the position of the ith student by 1 (i.e., moving
//     the ith student from position x to x + 1 or x - 1)
//
// Return the minimum number of moves required to move each student to a seat
// such that no two students are in the same seat.
//
// Note that there may be multiple seats or students in the same position at
// the beginning.

func minMovesToSeat(seats []int, students []int) int {
  slices.Sort(seats)
  slices.Sort(students)

  moves := 0
  for i, seat := range seats {
    moves += int(math.Abs(float64(students[i]) - float64(seat)))
  }

  return moves
}

func main() {
  function := minMovesToSeat
  seats := []int{}
  students := []int{}

  seats = []int{3,1,5}
  students = []int{2,7,4}
  fmt.Printf("Result  : %d\n", function(seats, students))
  fmt.Printf("Expected: %v\n\n", 4)

  seats = []int{4,1,5,9}
  students = []int{1,3,2,6}
  fmt.Printf("Result  : %d\n", function(seats, students))
  fmt.Printf("Expected: %v\n\n", 7)

  seats = []int{2,2,6,6}
  students = []int{1,3,2,6}
  fmt.Printf("Result  : %d\n", function(seats, students))
  fmt.Printf("Expected: %v\n\n", 4)
}

