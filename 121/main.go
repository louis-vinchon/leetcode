package main

import "fmt"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

// You are given an array prices where prices[i] is the price of a given stock
// on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock
// and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you
// cannot achieve any profit, return 0.

func maxProfit(prices []int) int {
  diff := 0

  minimumCandidate := prices[0]
  for i := 1 ; i < len(prices) ; i++ {
    price := prices[i]

    if price < minimumCandidate {
      minimumCandidate = price
    } else {
      candidateDiff := price - minimumCandidate
      if candidateDiff > diff {
        diff = candidateDiff
      }
    }
  }

  return diff
}

func main() {
  function := maxProfit
  prices := []int{}

  prices = []int{7,1,5,3,6,4}
  fmt.Printf("%v\n", function(prices)) // Expected: 5.

  prices = []int{7,6,5}
  fmt.Printf("%v\n", function(prices)) // Expected: 0.

  prices = []int{3,2,6,5,0,3}
  fmt.Printf("%v\n", function(prices)) // Expected: 4.

  prices = []int{10,12,11,15}
  fmt.Printf("%v\n", function(prices)) // Expected: 5.

  prices = []int{10,12,0,3}
  fmt.Printf("%v\n", function(prices)) // Expected: 3.
}

