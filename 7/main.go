package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/reverse-integer/

// Ok, face it, this is garbage. Complete garbage. Don't do that ever again.
// func reverse(x int) int {
//   inverse := 1
//   if x < 0 {
//     x = -x
//     inverse = -1
//   }
//
//   strNumb := fmt.Sprint(x)
//   runes := []rune(strNumb)
//   length := len(strNumb) 
//
//   for i := 0 ; i < length / 2 ; i++ {
//     temp := runes[i]
//     runes[i] = runes[length - 1 - i]
//     runes[length - 1 - i] = temp
//   }
//
//   res, ok := strconv.Atoi(string(runes))
//   if ok != nil {
//     fmt.Println("Not a number ?")
//     fmt.Println(fmt.Errorf(ok.Error()))
//   }
//
//   asInt32 := int(int32(res))
//   if asInt32 != res {
//     return 0
//   }
//
//   return inverse * asInt32
// }

// Much simpler without stirng conversions or what not:
// 1. pop the last sigit and insert it first into the reverse number (with modulo and division).
// 2. check for overflow
func reverse(x int) int {
  rev := 0

  // Do until we have extracted all digits from the number.
  for x != 0 {
    // Take out the last digit.
    pop := x % 10
    x /= 10

    // Shift the reverse left, and insert the digit.
    rev = rev * 10 + pop

    // Check for overflow.
    if rev > math.MaxInt32 || rev < math.MinInt32 {
      return 0
    }
  }

  return rev
}

func main() {
  var integer int = 154
  fmt.Printf("%v\n", reverse(integer))

  integer = -321
  fmt.Printf("%v\n", reverse(integer))

  integer = 1534236469
  fmt.Printf("%v\n", reverse(integer))
}

