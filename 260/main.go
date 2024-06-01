package main

import "fmt"

// https://leetcode.com/problems/single-number-iii/description/

// You can absolutely just count numbers in a map structure. And then only retain those that have a count of 1.
// Bbut there is a faster solution.

// Let's call A and B the two unique numbers.

func getBit(x int, k int) int {
  return (x >> k) & 1
}

func singleNumber(nums []int) []int {
  // XOR all numbers together. Number pairs will cancel each other. Leaving (A XOR B).
  xorAB := 0
  for _, num := range nums {
    xorAB ^= num
  }

  // Find the first bit set to 1 in (A XOR B). This is a bit that is present in
  // either A or B, but not both. We are guaranteed to get one because A and B
  // cannot be exactly the same number.
  pos := 0
  for getBit(xorAB, pos) == 0 {
    pos++
  }

  // XOR all numbers that have this same bit set to 1 with (A XOR B). Once
  // again, all number pairs will cancel each other, leaving only A or B (the
  // one of the two that actually has that bit set to one and not the other).
  A := 0
  for _, num := range nums {
    if getBit(num, pos) != 0 {
      A ^= num
    }
  }

  // XOR A with (A XOR B) to cancel A, leaving only B.
  B := xorAB ^ A

  return []int{A, B}
}

func main() {
  input := []int{ 1, 1, 2, 3, 5, 2 }
  fmt.Println(fmt.Sprintf("%o - %o", input, singleNumber(input)))

  input = []int{ -1, 0 }
  fmt.Println(fmt.Sprintf("%o - %o", input, singleNumber(input)))
}


