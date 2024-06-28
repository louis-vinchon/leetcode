package main

import "fmt"

// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/description/
//
// Given the root of a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus the sum of all keys greater than the original key in BST.
//
// As a reminder, a binary search tree is a tree that satisfies these constraints:
//
//     The left subtree of a node contains only nodes with keys less than the node's key.
//     The right subtree of a node contains only nodes with keys greater than the node's key.
//     Both the left and right subtrees must also be binary search trees.
//
// Constraints:
//     The number of nodes in the tree is in the range [1, 100].
//     0 <= Node.val <= 100
//     All the values in the tree are unique.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func bstToGstLeft(root *TreeNode, sum *int) *TreeNode {
  // Dig into the right.
  if root.Right != nil {
    bstToGstLeft(root.Right, sum)
  }

  // Add the sum of everything to the right (we start with sum = 0).
  root.Val += *sum
  // Update the sum.
  *sum = root.Val

  // Do the lfet branch. The sum will be added everytime.
  if root.Left != nil {
    bstToGstLeft(root.Left, sum)
  } else {
    // When going back up, we only save the value of the "leftest bottomest" branch, where the sum is greatest.
    *sum = root.Val
  }

  return root
}

func bstToGst(root *TreeNode) *TreeNode {
  incr := 0
  return bstToGstLeft(root, &incr)
}

func main() {
  function := bstToGst
  root := TreeNode{}

  root = TreeNode{
    Val: 3,
    Left: &TreeNode{
      Val: 1,
    },
    Right: &TreeNode{
      Val: 5,
    },
  }
  fmt.Printf("Result  : %v\n", function(&root))
  fmt.Printf("Expected: %v\n\n", root)
}

