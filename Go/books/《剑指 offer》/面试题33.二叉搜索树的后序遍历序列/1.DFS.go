package main

/**
解题思路：利用二叉搜索树特性：后序遍历：最后一个元素是根节点，从左到右第一个大于根节点的元素是右子树的根节点
时间复杂度：N^2
空间复杂度：N（递归的深度） 
 */
func verifyPostorder(postorder []int) bool {
	return recursion(postorder,0,len(postorder)-1)
}

// recursion 判断是否是二叉搜索树
func recursion(postorder []int,begin,end int) bool {
	// 退出条件
	if begin >= end {
		return true
	}
	
	cur := begin // cur:下标
	// 左子树的值 < 根节点（寻找右子树的根节点）
	for postorder[cur] < postorder[end] { // end 是根节点
		cur++ 
	}
	right := cur  // right:右子树的根节点下标
	// 右子树的值 > 根结点	
	for postorder[cur] > postorder[end] {
		cur++
	}
	return cur == end && // 若遍历到根节点，则说明是二叉搜索树
		recursion(postorder,begin,right-1) && // 判断左子树是否是二叉搜索树
		recursion(postorder,right,end-1) // 判断右子树是否是二叉搜索树
}

