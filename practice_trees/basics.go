package practice_trees

type TreeNode struct {
	Num   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(num int) *TreeNode {
	return &TreeNode{Num: num}
}

func PreOrder(Root *TreeNode, ans []int) []int {
	if Root != nil {
		ans = append(ans, Root.Num)
	}
	if Root.Left != nil {
		ans = PreOrder(Root.Left, ans)
	}
	if Root.Right != nil {
		ans = PreOrder(Root.Right, ans)
	}
	return ans
}
func InOrder(Root *TreeNode, ans []int) []int {
	if Root.Left != nil {
		ans = InOrder(Root.Left, ans)
	}
	if Root != nil {
		ans = append(ans, Root.Num)
	}
	if Root.Right != nil {
		ans = InOrder(Root.Right, ans)
	}
	return ans
}
func PostOrder(Root *TreeNode, ans []int) []int {
	if Root.Left != nil {
		ans = PostOrder(Root.Left, ans)
	}
	if Root.Right != nil {
		ans = PostOrder(Root.Right, ans)
	}
	if Root != nil {
		ans = append(ans, Root.Num)
	}
	return ans
}
