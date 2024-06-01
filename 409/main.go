package main

import "fmt"

// https://leetcode.com/problems/longest-palindrome/

// Given a string s which consists of lowercase or uppercase letters, return
// the length of the longest palindrome that can be built with those letters.
//
// Letters are case sensitive, for example, "Aa" is not considered a
// palindrome.

// The idea:
// - count all characters..
// - count how many pairs we can make AND, for odd length palyndromes, +1 for
//   the center character (if we have an odd number of a given character.
// func longestPalindrome(s string) int {
//   counts := make(map[rune]int)
//   for _, char := range s {
//     counts[char]++
//   }
//
//   res := 0
//   plusOne := 0
//   for _, count := range counts {
//     quotient := count >> 1
//     rest := count & 1
//
//     res += quotient << 1
//     if rest != 0 {
//       plusOne = 1
//     }
//   }
//
//   // fmt.Printf("res %d\n", res)
//   // fmt.Printf("plusOne %d\n", plusOne)
//
//   return res + plusOne
// }

// Optimized version where we do not use a hashmap.
// A to Z characters are "worth" (byte interpreted as digit) 65 to 90.
// a to z characters are worth 97 to 122.
// we don't actually care to remmeber that 'a' has 3 occurences, only that 'a
// character' has 3 occurences. As long as we don't mix the counts, we are good to go.
func longestPalindrome(s string) int {
  // The characters can only be letters and capital letters, hence 2 * 26.
  counts := make([]int, 52)

  // Notice how we can use both numbers and letters to do the math.
  for _, char := range s {
    // Is a capital letter. The index will start a 26.
    if char < 97 {
      counts[char - 65 + 26]++
      // counts[char - 'A' + 26]++
    } else { // Is a lowercase letter. Index will start at 0.
      counts[char - 97]++
      // counts[char - 'a']++
    }
  }
  // fmt.Printf("char %v\n", counts)

  res := 0
  plusOne := 0
  for _, count := range counts {
    // Rest of division by two.
    rest := count & 1

    if rest > 0 {
      plusOne = 1
    }

    res += count - rest
  }

  // fmt.Printf("res %d\n", res)
  // fmt.Printf("plusOne %d\n", plusOne)

  return res + plusOne
}

func main() {
  input := "axccccdd"
  fmt.Println(fmt.Sprintf("%d\n", longestPalindrome(input))) // Expected 7.

  input = "aabb"
  fmt.Println(fmt.Sprintf("%d\n", longestPalindrome(input))) // Expected 4.

  input = "abcd"
  fmt.Println(fmt.Sprintf("%d\n", longestPalindrome(input))) // Expected 1.


  // fmt.Println(fmt.Sprintf("%d", 'A'))
  // fmt.Println(fmt.Sprintf("%d", 'Z'))
  // fmt.Println(fmt.Sprintf("%d", 'a'))
  // fmt.Println(fmt.Sprintf("%d", 'z'))
}

