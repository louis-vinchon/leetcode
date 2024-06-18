package main

import (
	"fmt"
	"slices"
  "container/heap"
)

// https://leetcode.com/problems/most-profit-assigning-work/description/

// You have n jobs and m workers. You are given three arrays: difficulty,
// profit, and worker where:
//
//     difficulty[i] and profit[i] are the difficulty and the profit of the ith
//     job, and worker[j] is the ability of jth worker (i.e., the jth worker
//     can only complete a job with difficulty at most worker[j]).
//
// Every worker can be assigned at most one job, but one job can be completed
// multiple times.
//
//     For example, if three workers attempt the same job that pays $1, then
//     the total profit will be $3. If a worker cannot complete any job, their
//     profit is $0.
//
// Return the maximum profit we can achieve after assigning the workers to the
// jobs.
//
//
// Example 1:
// Input: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
// Output: 100
// Explanation: Workers are assigned jobs of difficulty [4,4,6,6] and they get a profit of [20,20,30,30] separately.
//
// Example 2:
// Input: difficulty = [85,47,57], profit = [24,66,99], worker = [40,25,25]
// Output: 0
//
//
// Constraints:
//     n == difficulty.length
//     n == profit.length
//     m == worker.length
//     1 <= n, m <= 10^4
//     1 <= difficulty[i], profit[i], worker[i] <= 10^5

type Job struct {
  profit int
  difficulty int
}

type PriorityQueue []*Job

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
  // We want Pop to give us the highest, not lowest, priority so we use greater than here.
  return pq[i].profit > pq[j].profit
}

func (pq PriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
  item := x.(*Job)
  *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
  old := *pq
  n := len(old)
  item := old[n-1]
  old[n-1] = nil  // avoid memory leak
  *pq = old[0 : n-1]
  return item
}

// Doesn't work.
func (pq PriorityQueue) String() {
  for _, item := range pq {
    fmt.Printf("%v ", *item)
  }
  fmt.Printf("\n")
}



func maxProfitAssignment(difficulty []int, profits []int, workers []int) int {
  jobs := []Job{}
  for i := 0 ; i < len(difficulty) ; i++ {
    jobs = append(jobs, Job{
      difficulty: difficulty[i],
      profit: profits[i],
    })
  }

  slices.SortFunc(jobs, func(jobA, jobB Job) int {
    return jobA.difficulty - jobB.difficulty
  })
  fmt.Printf("jobs by growing difficulty %v\n", jobs)

  maxProfit := 0
  slices.Sort(workers)
  pq := make(PriorityQueue, 0)
  index := 0
  for _, worker := range workers {
    fmt.Printf("---\nworker %v\n", worker)
    for index < len(jobs) {
      fmt.Printf("job %d %v\n", index, jobs[index])
      if jobs[index].difficulty <= worker {
        fmt.Printf("push\n")
        heap.Push(&pq, &(jobs[index]))
        index++
      } else {
        fmt.Printf("break\n")
        break
      }
    }

    fmt.Printf("pq %v\n", pq)

    if len(pq) > 0 {
      fmt.Printf("best in pq %v\n", pq[0])
      maxProfit += pq[0].profit
    }
  }

  return maxProfit
}

func main() {
  function := maxProfitAssignment
  difficulty := []int{}
  profit := []int{}
  worker := []int{}

  difficulty = []int{2,4,6,8,10}
  profit = []int{10,20,30,40,50}
  worker = []int{4,5,6,7}
  fmt.Printf("Result  : %v\n", function(difficulty, profit, worker))
  fmt.Printf("Expected: %v\n", 100)


  difficulty = []int{85,47,57}
  profit = []int{24,66,99}
  worker = []int{40,25,25}
  fmt.Printf("Result  : %v\n", function(difficulty, profit, worker))
  fmt.Printf("Expected: %v\n", 0)
}

