package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/balanced-binary-tree/description/

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}


func isBalanced(root *TreeNode) bool {
  balanced := true

  // You need to declare the type separately for recursive functions.
  var getLength func(node *TreeNode) int
  getLength = func(node *TreeNode) int {
    if !balanced {
      return 0
    }
    if node == nil {
      return 0
    }

    leftLength := getLength(node.Left)
    rightLength := getLength(node.Right)

    if math.Abs(float64(leftLength) - float64(rightLength)) > 1 {
      balanced = false
      return 0
    }

    // fmt.Printf("%v %t\n", 1 + max(leftLength, rightLength), true)

    return 1 + max(leftLength, rightLength)
  }

  getLength(root)

  return balanced
}

func main() {
  function := isBalanced
  root := TreeNode{}

  // This would take a long ass time without caching the results.
  root = TreeNode{
    Right: &TreeNode{}, // 1
    Left: &TreeNode{
      Right: &TreeNode{}, // 2
      Left: &TreeNode{
        Right: &TreeNode{}, // 3
        Left: &TreeNode{}, // 3
      },
    },
  }
  fmt.Printf("Result   %v\n", function(&root))
  fmt.Printf("Expected %v\n\n", false)
}

