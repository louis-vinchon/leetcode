package main

import "fmt"

// https://leetcode.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/description/

// Read carefully:
// i < j <= k
// a = arr[i] ^ ... ^ a[j -1]
// b = arr[j] ^ ... ^ a[k]

// This is slow af.
// func countTriplets(arr []int) int {
//   matches := 0
//   for i := 0 ; i < len(arr) - 1 ; i++ {
//     for k := i + 1 ; k < len(arr); k++ {
//       for j := i + 1 ; j <= k ; j++ {
//         a, b := 0, 0
//
//         for x := i ; x < j ; x++ {
//           a ^= arr[x]
//         }
//         for y := j ; y <= k ; y++ {
//           b ^= arr[y]
//         }
//
//         if a == b {
//           matches++
//         }
//       }
//     }
//   }
//
//   return matches
// }


// Optimisation 1:
// a = arr[i] ^ ... ^ a[j -1]
// b = arr[j] ^ ... ^ a[k]
// a == b is equivalent to a ^ b == 0, which is basically arr[i] ^ ... ^ arr[k] == 0, so the position of j doesn't matter.
// Meaning every possible position of j in the [i:k] range is a valid answer.
// So we only need to iterate i and k.

// Optimisation 2:
// We build the XOR iteratively and check every loop for k.
// Thanks to the first optimization this also gets read of the 4th level loops.
func countTriplets(arr []int) int {
  matches := 0

  for i, val := range arr {
    for k := i + 1 ; k < len(arr) ; k++ {
      val ^= arr[k]
      if val == 0 {
        matches += k - i
      }
    }
  }

  return matches
}

func main() {
  input := []int{ 1, 1, 1, 1, 1 }
  fmt.Println(fmt.Sprintf("%o - matches: %d", input, countTriplets(input)))

  // input = []int{ 2,3,1,6,7 }
  // fmt.Println(fmt.Sprintf("%o - matches: %d", input, countTriplets(input)))
}

