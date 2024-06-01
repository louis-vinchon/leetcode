
package main

import (
	"fmt"
  "math"
)

// https://leetcode.com/problems/reverse-integer/

func divide(dividend int, divisor int) int {
  absDive := int(math.Abs(float64(dividend)))
  absDivo := int(math.Abs(float64(divisor)))

  quotient := 0
  if absDivo == 1 { // Save some time because why not.
    quotient = absDive
  } else {
    for absDive >= absDivo {
      pumpedUpDivisor := absDivo
      quotientIncrement := 1

      // So, normally you would just remove divisor from dividend. Over and over again.
      // To save time, you could multiplicate your divisor to get close to the devidend.
      // No multiplication? Bitshift to the rescue!
      //
      // So we basically "multiply" (bitshift) our divisor by 2 until it's at most half of the dividend (one bit place smaller).
      // Everytime we multiply our divisor by two means we are going to add twice as much to the quotient.
      for absDive >= pumpedUpDivisor << 1 {
        pumpedUpDivisor <<= 1
        quotientIncrement <<= 1
      }

      absDive -= pumpedUpDivisor
      quotient += quotientIncrement
    }
  }

  negative := (dividend < 0) != (divisor < 0)

  if negative {
    quotient = -quotient
  }

  if quotient > math.MaxInt32 {
    return math.MaxInt32
  } else if quotient < math.MinInt32 {
    return math.MinInt32
  }

  return quotient
}

func main() {
  dividend, divisor := -1, 1
  fmt.Printf("%v\n\n", divide(dividend, divisor))

  // dividend, divisor = -2147483648, -3
  // fmt.Printf("%v\n\n", divide(dividend, divisor))
}

