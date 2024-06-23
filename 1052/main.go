package main

import "fmt"

// https://leetcode.com/problems/grumpy-bookstore-owner/description/

// There is a bookstore owner that has a store open for n minutes. Every minute, some number of customers enter the store. You are given an integer array customers of length n where customers[i] is the number of the customer that enters the store at the start of the ith minute and all those customers leave after the end of that minute.
//
// On some minutes, the bookstore owner is grumpy. You are given a binary array grumpy where grumpy[i] is 1 if the bookstore owner is grumpy during the ith minute, and is 0 otherwise.
//
// When the bookstore owner is grumpy, the customers of that minute are not satisfied, otherwise, they are satisfied.
//
// The bookstore owner knows a secret technique to keep themselves not grumpy for minutes consecutive minutes, but can only use it once.
//
// Return the maximum number of customers that can be satisfied throughout the day.
//
// Constraints:
//     n == customers.length == grumpy.length
//     1 <= minutes <= n <= 2 * 10^4
//     0 <= customers[i] <= 1000
//     grumpy[i] is either 0 or 1.


func maxSatisfied(customers []int, grumpy []int, minutes int) int {
  notGrumpySatisfied := 0
  maxGrumpy := 0
  for i := 0 ; i < minutes ; i++ {
    maxGrumpy += customers[i] * grumpy[i]
    notGrumpySatisfied += customers[i] * (1-grumpy[i])
  }

  grumpyPotential := maxGrumpy
  for i := minutes ; i < len(customers) ; i++ {
    grumpyPotential -= customers[i-minutes] * grumpy[i-minutes]
    grumpyPotential += customers[i] * grumpy[i]
    maxGrumpy = max(maxGrumpy, grumpyPotential)

    notGrumpySatisfied += customers[i] * (1-grumpy[i])
  }

  return maxGrumpy + notGrumpySatisfied
}

func main() {
  function := maxSatisfied
  customers := []int{}
  grumpy := []int{}
  minutes := 0

  customers = []int{1,0,1,2,1,1,7,5}
  grumpy =    []int{0,1,0,1,0,1,0,1}
  minutes = 3
  fmt.Printf("Result  : %v\n", function(customers, grumpy, minutes))
  fmt.Printf("Expected: %v\n", 16)

  customers = []int{1}
  grumpy =    []int{0}
  minutes = 1
  fmt.Printf("Result  : %v\n", function(customers, grumpy, minutes))
  fmt.Printf("Expected: %v\n", 1)

  customers = []int{1,1,1,1,1,1,1}
  grumpy =    []int{1,1,1,1,1,1,1}
  minutes = 4
  fmt.Printf("Result  : %v\n", function(customers, grumpy, minutes))
  fmt.Printf("Expected: %v\n", minutes)

  customers = []int{1,1,1,1,1,1,1}
  grumpy =    []int{1,0,1,0,0,1,1}
  minutes = 3
  fmt.Printf("Result  : %v\n", function(customers, grumpy, minutes))
  fmt.Printf("Expected: %v\n", 5)
}


