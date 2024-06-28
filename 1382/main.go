package main

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/balance-a-binary-search-tree/description/

// Given the root of a binary search tree, return a balanced binary search tree
// with the same node values. If there is more than one answer, return any of
// them.
//
// A binary search tree is balanced if the depth of the two subtrees of every
// node never differs by more than 1.
//
// Constraints:
//   The number of nodes in the tree is in the range [1, 10^4].
//   1 <= Node.val <= 10^5


type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func getSortedValues(node *TreeNode) []int {
  values := []int{}
  if node.Left != nil {
    values = slices.Concat(values, getSortedValues(node.Left))
  }

  values = append(values, node.Val)

  if node.Right != nil {
    values = slices.Concat(values, getSortedValues(node.Right))
  }

  return values
}

func makeBalancedTree(values []int) *TreeNode {
  fmt.Printf("list %v\n", values)
  midIndex := len(values) / 2
  // len == 3  ->   1
  // len == 2  ->   1
  // len == 1  ->   0

  node := &TreeNode{
    Val: values[midIndex],
  }

  if midIndex > 0 {
    leftValues := values[:midIndex-1]
    leftTree := makeBalancedTree(leftValues)
    node.Left = leftTree
  }

  if midIndex < len(values) - 1 {
    rightValues := values[midIndex+1:]
    rightTree := makeBalancedTree(rightValues)
    node.Right = rightTree
  }

  return node
}

func balanceBST(root *TreeNode) *TreeNode {
  values := getSortedValues(root)

  fmt.Printf("values %v\n", values)

  balancedTree := makeBalancedTree(values)

  return balancedTree
}

func main() {
  function := balanceBST
  root := &TreeNode{}

  fmt.Printf("Result  : %v\n", function(root))
}

