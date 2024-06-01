package main

import (
  "fmt"
)

// https://leetcode.com/problems/house-robber/

// You are a professional robber planning to rob houses along a street. Each
// house has a certain amount of money stashed, the only constraint stopping
// you from robbing each of them is that adjacent houses have security systems
// connected and it will automatically contact the police if two adjacent
// houses were broken into on the same night.
//
// Given an integer array nums representing the amount of money of each house,
// return the maximum amount of money you can rob tonight without alerting the
// police.

// DP problem apparently. I can't fucking discern a "sub problem".
// Apparently. It's about "3rd house + 1st house" to the "second house" by itself ??
// These become agregates overtime but I just don't get it. Why Are we
// respecting the spacing backwards, but never worried about the spacing
// forwards?????
// I hate this problem. It works. But it makes no sense to me.
func rob(nums []int) int {
  fmt.Printf("NUMS %d\n", nums)
  firstMove := 0
  secondMove := 0
  for i := 0; i < len(nums); i++{
    t := max(nums[i] + firstMove, secondMove) // So. If we translate this: if at some point, by adding the current house with the former-former one(aggregated), we can't beat the former house, then we "skip", but the score of the current house becomes the score of the former house? i.e If (house0 + house2) can't beat (house1), then the house(1) score is optimal up to (house2), regardless of what comes next???
    fmt.Printf("first move %d\n", firstMove)
    fmt.Printf("second move %d\n", secondMove)
    fmt.Printf("value %d\n", nums[i])
    fmt.Printf("current move %d\n", t)
    fmt.Printf("----------\n")
    firstMove = secondMove
    secondMove = t
  }
  return secondMove
}

func main() {
  input := []int{ 1, 9, 1, 9, 1 }
  fmt.Printf("%d\n", rob(input)) // Expected 18.

  input = []int{ 1, 9, 1, 1, 9, 1 }
  fmt.Printf("%d\n", rob(input)) // Expected 18.

  input = []int{ 1, 9, 1, 1, 1, 9, 1 }
  fmt.Printf("%d\n", rob(input)) // Expected 19.
}

