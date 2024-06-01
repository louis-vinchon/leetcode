package main

import "fmt"

// https://leetcode.com/problems/find-common-characters/description/

// Given a string array words, return an array of all characters that show up
// in all strings within the words (including duplicates). You may return the
// answer in any order.

func debugMatrix(matrix [][]uint8) {
  for _, row := range matrix {
    fmt.Printf("%v\n", row)
  }
}

// words are only composed of lowercase characters
func commonCharsV1(words []string) []string {
  // One row per word, and each row is an array to count each letter. uint8
  // becaue we have at most 100 letters in a word, and uint8 goes until 256.
  // Can we use int? Yeah we can. It's surprisingly slower though.
  counts := make([][]uint8, len(words))

  for i, word := range words {
    counts[i] = make([]uint8, 26)

    for _, char := range word {
      // Usual trick to get a uint value from a rune (aka byte).
      counts[i][char - 'a']++
    }
  }

  debugMatrix(counts)

  duplicateChars := []string{}
  for i := 0 ; i < 26 ; i++ {
    // There are maximum 100 letters in a word.
    min := uint8(100)
    for _, word := range counts {
      if word[i] < min {
        min = word[i]
      }
    }
    // fmt.Printf("letter %q - min: %d\n", 'a' + i, min)

    for j := uint8(0) ; j < min ; j++ {
      duplicateChars = append(duplicateChars, string(rune('a' + i)))
      // fmt.Printf("add %q\n", 'a' + i)
    }
  }

  return duplicateChars
}

var commonChars = commonCharsV1

func main() {
  words := []string{ "asd", "bfa" }
  fmt.Printf("res %s\n", commonChars(words)) // Expected: ['a'].

  words = []string{ "bella", "label", "roller" }
  fmt.Printf("res %s\n", commonChars(words)) // Expected: ['e', 'l', 'l'].
}

