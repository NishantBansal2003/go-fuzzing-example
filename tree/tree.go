package tree

// TreeNode represents a node in a binary tree with integer values.
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

// BuildTree creates a complete binary tree of the specified depth.
// Each node's Value is set to the current depth. Returns nil if the depth is
// not positive.
func BuildTree(depth int) *TreeNode {
	if depth <= 0 {
		return nil
	}
	return &TreeNode{
		Left:  BuildTree(depth - 1),
		Right: BuildTree(depth - 1),
		Value: depth,
	}
}
