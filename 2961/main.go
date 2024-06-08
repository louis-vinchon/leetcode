package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/double-modular-exponentiation/description/



// (A^B)%6
// ((6x + b)*(6x + b))^(B-1)    % 6
// (36x^2 + 12xb + b^2)^(B-1)    % 6
// (36x^2 + 12xb + b^2)*(6x + b)^(B-2)    % 6
// (6s ish + b^3)^(B-2)    % 6
// ...
// (A^B)%6 == (A%6)^B
//
// So, as weird as it sounds, if you followed the little math I am capable of above.
// We don't need to do the exponent on the whole number, just its modulo. Same result.
// What you'll also notice is that the modulo value, while it can vary, always loops back.
// We mesure this loopback and use it to decimate the power value.
//
// This calculates (numb^power)%dividend in a VERY efficient fashion.
func powModulo(numb, power, dividend int) int {
  if dividend == 1 { return 0 }

  // if power == 0 { return 1 } // Not possible acordng to the problem.
  if power == 1 { return numb % dividend }

  powedNumb := numb % dividend
  if powedNumb < 2 { return powedNumb } // If it's 0 or one, 0^whatever = 0, and 1^whatever = 1.

  power--

  initial := powedNumb
  loopbackLength := 0
  for power > 0 {
    power--
    loopbackLength++

    powedNumb *= numb
    powedNumb %= dividend

    if powedNumb == initial {
      power %= loopbackLength
      loopbackLength = 0
    }
  }

  return powedNumb
}

// Brute force solution. I won't even try it given that the numbers range
// 1-1000, I let you decide what is going to happen if you do 1000^1000 (10e3000).
func getGoodIndicesBruteForce(variables [][]int, target int) []int {
  goodIndices := []int{}
  for i, row := range variables {
    score := int(math.Pow(float64(row[0]), float64(row[1])))
    score %= 10
    score = int(math.Pow(float64(score), float64(row[2])))
    score %= row[3]

    if score == target {
      goodIndices = append(goodIndices, i)
    }
  }

  return goodIndices
}

func getGoodIndicesOpti(variables [][]int, target int) []int {
  goodIndices := []int{}
  for i, row := range variables {
    score := powModulo(row[0], row[1], 10)
    score = powModulo(score, row[2], row[3])

    if score == target {
      goodIndices = append(goodIndices, i)
    }
  }

  return goodIndices
}

func main() {
  function := getGoodIndicesOpti
  variables := [][]int{}
  target := 0

  variables = [][]int{{2,3,3,10},{3,3,3,1},{6,1,1,4}}
  target = 2
  fmt.Printf("%v\n", function(variables, target)) // Expected: [0 2]

  variables = [][]int{{39,3,1000,1000}}
  target = 17
  fmt.Printf("%v\n", function(variables, target)) // Expected: []

  // fmt.Printf("inputs %d %d %d\n", 2, 3, 6)
  // fmt.Printf("modulo %v\n\n", powModulo(2,3,6))
  // fmt.Printf("inputs %d %d %d\n", 2, 1231738, 6)
  // fmt.Printf("modulo %v\n", powModulo(2,123173,6))

  // fmt.Printf("\nmodulo %v\n\n", 2 % 6)
  // fmt.Printf("modulo %v\n\n", 4 % 6)
  // fmt.Printf("modulo %v\n\n", 8 % 6)
  // fmt.Printf("modulo %v\n\n", 16 % 6)
  // fmt.Printf("modulo %v\n\n", 32 % 6)
  // fmt.Printf("modulo %v\n\n", 64 % 6)
  //
  // fmt.Printf("\nmodulo %v\n\n", 3 % 6)
  // fmt.Printf("modulo %v\n\n", 3*3 % 6)
  // fmt.Printf("modulo %v\n\n", 3*3*3 % 6)
  // fmt.Printf("-- %v\n\n", powModulo(3,3,6))
  //
  // fmt.Printf("\nmodulo %v\n", 3 % 8)
  // fmt.Printf("modulo %v\n", 3*3 % 8)
  // fmt.Printf("modulo %v\n", 3*3*3 % 8)
  // fmt.Printf("modulo %v\n", 3*3*3*3 % 8)
  // fmt.Printf("modulo %v\n", 3*3*3*3*3 % 8)
  // fmt.Printf("-- %v\n\n", powModulo(3,4,8))
  //
  // fmt.Printf("modulo %v\n", 2 % 10)
  // fmt.Printf("modulo %v\n", 2*2 % 10)
  // fmt.Printf("modulo %v\n", 2*2*2 % 10)
  // fmt.Printf("modulo %v\n", 2*2*2*2 % 10)
  // fmt.Printf("-- %v\n\n", powModulo(2,4,10))
  //
  // fmt.Printf("modulo %v\n", int(math.Pow(10,12)) % 6)
  // fmt.Printf("-- %v\n\n", powModulo(10,12,6))
}

