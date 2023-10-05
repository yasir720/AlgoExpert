package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    
    ans := [][]int{}
    queue := []*TreeNode{ root }
    
    for len(queue) > 0 {
        tmp := []int{}
        
        for _, node := range queue {
            queue = queue[1:]
            tmp = append(tmp, node.Val)
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        ans = append(ans, tmp)
    }
    
    return ans
}
