package main

import "fmt"

// https://leetcode.com/problems/climbing-stairs/description/

// The idea is that, while it's unintuitive to know how many ways there are to
// solve the problem when many steps are left, it is however quite easy to do
// so when there are very few left.
//
// So we use recursion and count "backwards".
//
// Note: this version probably isn't the fastest, but it requires no math
// knowledge and would work for any step sizes.
//
// it is a "Dynamic programming" problem where we are likely to solve the same
// "intermediate problems" several times. Hence we maintain a "cache" for
// intermediate solutions to those problems (it saves a lot of time).

var steps = []int{1,2}
var solutionsCache = make(map[int]int)

func climbStairs(n int) int {
  val, ok := solutionsCache[n]
  if ok {
    return val
  }

  result := 0
  for _, step := range steps {
    remainingSteps := n - step

    if remainingSteps < 0 {
      // The step is too big, this is not an acceptable solution.
      break
    }
    if remainingSteps == 0 {
      // The step fits perfectly. This is a solution.
      result += 1
    } else {
      // We are not at a solution yet, keep going deeper.
      result += climbStairs(remainingSteps)
    }
  }

  solutionsCache[n] = result

  return result
}

func main() {
  // This would take a long ass time without caching the results.
  input := 45
  fmt.Println(fmt.Sprintf("%d steps - %d", input, climbStairs(45)))
}

