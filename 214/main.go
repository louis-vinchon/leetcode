package main

import (
	"fmt"
	"slices"
	"strings"
)

// https://leetcode.com/problems/shortest-palindrome/description/

// You are given a string s. You can convert s to a
// palindrome
// by adding characters in front of it.
//
// Return the shortest palindrome you can find by performing this transformation.

// My idea: we reverse the string, and then we compare to the original. We
// slide them apart one character until there is a match. Then we cut out
// however many characters we had to slide out. And we are done. It's too slow
// apparently.
func shortestPalindromeV1(s string) string {
  reverse := []rune(s)
  slices.Reverse(reverse)

  offset := 0
  for ; offset < len(s) ; offset++ {
    j := 0
    remainingChars := len(reverse) - offset
    // fmt.Printf("%v\n", strings.Repeat(" ", offset) + s)
    // fmt.Printf("%v\n", string(reverse))
    for j < remainingChars {
      if reverse[offset + j] != []rune(s)[j] {
        break
      }

      j++
    }

    if j == remainingChars {
      break
    }
  }

  return string(reverse[:offset]) + s
}

// Fancy, more efficient solution (KMP).

// https://cp-algorithms.com/string/prefix-function.html
// Brute force LPS table algorithm O(n^3).
// Careful, in the article they use str.substr(pos, length), in contrast to
// str[lowerBound:upperBound].
func makeLpsTableV1(pattern string) []int {
  n := len(pattern)
  lps := make([]int, n)

  // For each sub-pattern length.
  for i := 1 ; i < n ; i++ {
    fmt.Printf("=== i %d ======\n", i)
    // For each possible sub-pattern for the given length.
    for k := 1 ; k <= i ; k++ {
      // fmt.Printf("--\npattern %s\n", pattern)
      // fmt.Printf("subpat1 %v\n", pattern[:k])
      // fmt.Printf("subpat2 %v\n", strings.Repeat(" ", i-k+1) + pattern[i-k+1:i+1])

      if pattern[:k] == pattern[i-k+1:i+1] {
        // Reminder: it's not about counting, it's about finding the longest prefix possible.
        // It's in the name after all: Longest prefix suffix.
        lps[i] = k
      }
    }
    // fmt.Printf("=== i %d | lps %d ======\n", i, lps[i])
  }

  return lps
}

// https://cp-algorithms.com/string/prefix-function.html
// Video for step 4. WATCH IT.
// https://www.youtube.com/watch?v=jXUP-sAzXRQ
// This is incomprehensible to me.
// STEP 1
// a b a b x x a b a b a...
//                     ^
//                     i
//
// 0 0 1 2 0 0 1 2 3 4 ?
//                   ^
//         j = lps[i-1]
//
// STEP 3
// a b a b x x a b a b a
// ======= ^     =======
// pattern[j]
//
// So far:
// - we got the longest prefix length for i-1: lps[i-1] == 4.
// - underlined above is where the longest prefix AND suffix is for i-1.
// - we know our suffix prefix can be at most 1 longer.
// - So:
//   - The natural continuation of the previous longest prefix is s[4] (s[lps[i-1]])
//   - And the natural continuation of the previous longest suffix is s[i] (s[i-1+1]).
//
// If they are the same, then congratulations, you can just do previous longest + 1 and be done with i.
//
// However, if they are not the same. Follow below:
//
// STEP 4
// 0 0 1 2 0 0 1 2 3 4 ?
//                   ^
//         j = lps[i-1] = 4
// a b a b x x a b a b a
// ======= ^     =======
//
// We need to find the next lps item candidate. How so? We only have one value
// for lps[i-1] after all.
// The trick is: the underlined parts are the same string and a substring of
// the whole pattern, right? We have already calculated the lps for that
// substring at some point, right? Well, any prefix suffix to that substring,
// will ALSO be a prefix suffix for the entire pattern. See below.
//
// Finding out what the lsp for the substring is:
// =======
// 0 0 1 2 0 0 1 2 3 4 ?
//       ^
//  j = lps[j-1] = lps[4-1] = 2
//
// a b a b x x a b a b a
// =======       ======= (first lps check)
// ~~~ ~~~       ~~~ ~~~ (second lps check)
//     ^               ^
//  new j (2)          this is the interesting part
//
// So, to summarize things quickly:
// - we tried to see, with the first lps lookup, if we could easily extend the previous lps. But it failed.
// - we can lookup the lps for the substring itself.
//   This new prefix suffix is, as you can see, a suffix prefix for the entire pattern as well.
//   Thus we can try to see if we can extend it by one just like before. And if
//   it fails, recursively do what we just did until we reach lsp = 0. At which
//   point we basically are out of luck and starting counting from 0 again. No
//   extension is possible at all.
//
// The LPS table is secretly recursive!
//
// It is still wizardry to me how someone can figure out this kind of stuff.
// But it works, and is now understandable to me.
func makeLpsTableOptimized(pattern string) []int {
  n := len(pattern)
  lps := make([]int, len(pattern))

  for i := 1; i < n; i++ {
    j := lps[i-1];
    for j > 0 && pattern[i] != pattern[j] {
      j = lps[j-1];
    }
    if (pattern[i] == pattern[j]) {
      j++;
    }
    lps[i] = j;
  }

  return lps
}

// I do not understand the logic here.
func shortestPalindromeV2(s string) string {
  if len(s) == 0 {
    return s
  }

  reverse := []rune(s)
  slices.Reverse(reverse)

  // Additional delimiter to not "mix" the two strings.
  sPlusRev := string(reverse) + "#" + s
  fmt.Printf("%s\n", sPlusRev)

  lps := makeLpsTableOptimized(sPlusRev)

  longestPrefixForFullLength := lps[len(sPlusRev) - 1]
  fmt.Printf("%d\n", longestPrefixForFullLength)
  // This will be the longest prefix / reversed suffix. 

  // fmt.Printf("%v\n", lps)

  return string(reverse[:len(s) - longestPrefixForFullLength + 1]) + s
}

func main() {
  shortestPalindrome := shortestPalindromeV2
  input := "aacecaaa"
  output := shortestPalindrome(input)
  fmt.Printf("%s\n", strings.Repeat(" ", len(output) - len(input)) + input) // Expected: "aaacecaaa".
  fmt.Printf("%s\n", output) // Expected: "aaacecaaa".

  // input = "asa"
  // fmt.Printf("%s\n", shortestPalindrome(input)) // Expected: "asa".
  //
  // input = "abcd"
  // fmt.Printf("%s\n", shortestPalindrome(input)) // Expected: "dcbabcd".
  //
  // input = ""
  // fmt.Printf("%s\n", shortestPalindrome(input)) // Expected: "dcbabcd".
}

