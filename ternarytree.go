package ternarytree

// TernaryTree is a ternary search tree data structure.
type TernaryTree struct {
	head     *treeNode
	hasEmpty bool
}

type treeNode struct {
	char  byte
	loKid *treeNode
	eqKid *treeNode
	hiKid *treeNode
	data  *string
}
