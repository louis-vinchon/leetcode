package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/maximum-total-importance-of-roads/description/

// You are given an integer n denoting the number of cities in a country. The
// cities are numbered from 0 to n - 1.
//
// You are also given a 2D integer array roads where roads[i] = [ai, bi]
// denotes that there exists a bidirectional road connecting cities ai and bi.
//
// You need to assign each city with an integer value from 1 to n, where each
// value can only be used once. The importance of a road is then defined as the
// sum of the values of the two cities it connects.
//
// Return the maximum total importance of all roads possible after assigning
// the values optimally.
//
// Constraints:
//     2 <= n <= 5 * 10^4
//     1 <= roads.length <= 5 * 10^4
//     roads[i].length == 2
//     0 <= ai, bi <= n - 1
//     ai != bi
//     There are no duplicate roads.

type City struct {
  occurences int
  index      int // The index of the item in the heap.
}

func (city City) String() string {
  return fmt.Sprintf("{ occurences: %d, index: %d }", city.occurences, city.index)
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*City

// Note: this DOES NOT print the priority queue in its actual order.
func (pq PriorityQueue) String() string {
  str := "{\n"
  for _, city := range pq {
    str += fmt.Sprintf("  %v\n", *city)
  }
  str += "}\n"
  return str
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, k int) bool {
  return pq[i].occurences > pq[k].occurences
}

func (pq PriorityQueue) Swap(i, k int) {
  pq[i], pq[k] = pq[k], pq[i]
  pq[i].index = i
  pq[k].index = k
}

func (pq *PriorityQueue) Push(x any) {
  n := len(*pq)
  item := x.(*City)
  item.index = n
  *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
  old := *pq
  n := len(old)
  item := old[n-1]
  old[n-1] = nil  // avoid memory leak
  item.index = -1 // for safety
  *pq = old[0 : n-1]
  return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) increment(item *City) {
  item.occurences++
  heap.Fix(pq, item.index)
}

func maximumImportance(n int, roads [][]int) int64 {
  // Ranked list of *City objects. We rank cities based on the number of roads connecting them.
  priorityQueue := PriorityQueue{}

  // Map city number to *City object. This is so that we can quickly access each city object by their city number.
  index := make(map[int]*City)

  for _, road := range roads {
    for _, cityNum := range road {
      city, ok := index[cityNum]
      if !ok {
        newCity := City{
          occurences: 1,
        }

        index[cityNum] = &newCity
        heap.Push(&priorityQueue, &newCity)
      } else {
        priorityQueue.increment(city)
      }
    }
  }

  // The cities with the most roads have the best score.
  var result int64 = 0
  for len(priorityQueue) > 0 {
    item := heap.Pop(&priorityQueue).(*City)
    result += int64(item.occurences * n)
    n--
  }

  return result
}

func main() {
  function := maximumImportance
  n := 0
  roads := [][]int{}

  n = 5
  roads = [][]int{{0,1},{1,2},{2,3},{0,2},{1,3},{2,4}}
  fmt.Printf("Result  : %v\n", function(n, roads))
  fmt.Printf("Expected: %v\n\n", 43)

  n = 5
  roads = [][]int{{0,3},{2,4},{1,3}}
  fmt.Printf("Result  : %v\n", function(n, roads))
  fmt.Printf("Expected: %v\n\n", 20)
}

