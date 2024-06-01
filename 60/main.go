package main

import (
  "fmt"
  "math"
  "strconv"
)

// https://leetcode.com/problems/permutation-sequence/


func getPermutation(n int, k int) string {
  fmt.Printf("n = %d - k = %d\n", n, k)
  permutations := 1
  for i := 2 ; i <= n ; i++ {
    permutations *= i
  }

  // This is an integer whose bits we are going to use to count which numbers have been used in the permutation or not.
  // A bit set to one means the number is available, 0 means it's been taken.
  // Since n <= 9, we are only ever going to use 9 of the bits.
  availableNumbers := math.MaxInt16

  // End result.
  res := ""

  nDesc := n
  for len(res) < n {
    // Find out how many permutations start with each number. It is the total
    // number of permuatations divided by how many numbers there are. e.g. n =
    // 3 implies 2 permutations starting with 1, 2 with 2, 2 with 3. Incidently
    // this is (n - 1)!
    permutations /= nDesc
    nDesc--
    fmt.Printf("permutations per number %d\n", permutations)

    // Find which number place we should use.
    // e.g. n = 3 and k = 3 means we use the third permutation out of 6, so our
    // permutation should start with the second number (place is 2).
    place := int(math.Ceil(float64(k) / float64(permutations)))
    fmt.Printf("place %v\n", place)

    // Only keep the rest of k.
    k %= permutations
    if k == 0 {
      k = permutations
    }

    // Now we find out which number we can actually use. The second number
    // might have been taken already, so we'll check all numbers in succession
    // until one is free.
    takeNumber := 0
    i := 1
    for takeNumber == 0 {
      fmt.Printf("i %d\n", i)
      // Mask to check individual bits. We don't care about not starting at 0,
      // because we only ever have n <= 9, so we only need 9 places out of 16.
      mask := 1 << i
      fmt.Printf("available numbers %016b\n", availableNumbers)
      fmt.Printf("mask              %016b\n", mask)

      // If number is not available.
      if availableNumbers & mask == 0 {
        // Next loop / next number.
      } else {
        // Count one towards the number place.
        place--
        if place == 0 {
          // If we have reached the desired number, pick it.
          takeNumber = i
          // Flip the bit for the taken number (to 0).
          availableNumbers ^= mask

          fmt.Printf("take number %d\n", takeNumber)
          fmt.Printf("remaining available numbers %#b\n", availableNumbers)
        }
      }
      i++
    }

    res = res + strconv.Itoa(takeNumber)

    fmt.Printf("res: %s\n\n", res)
  }

  return res
}

func main() {
  n, k := 3, 2
  fmt.Printf("%v\n\n", getPermutation(n, k)) // Should be: 132
  // 123
  // 132 <
  // 213
  // 231
  // 312
  // 321

  // n, k = 3, 2
  // fmt.Printf("%v\n\n", getPermutation(n, k)) // Should be: 213
  //
  // n, k = 4, 9
  // fmt.Printf("%v\n\n", getPermutation(n, k)) // Should be: 2314

  // n, k = 3, 1
  // fmt.Printf("%v\n\n", getPermutation(n, k)) // Should be: 123
}

