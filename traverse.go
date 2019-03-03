package ternarytree

// Traverse recursively applies a function to all stored strings.
func (tree *TernaryTree) Traverse(f func(string)) {
	if tree.hasEmpty {
		f("")
	}
	traverseVisitor(tree.head, f)
}

func traverseVisitor(node *treeNode, f func(string)) {
	if node == nil {
		return
	}
	traverseVisitor(node.loKid, f)
	traverseVisitor(node.eqKid, f)
	if node.data != nil {
		f(*node.data)
	}
	traverseVisitor(node.hiKid, f)

}

// Sort returns a sorted list of strings contained in the tree.
func (tree *TernaryTree) Sort() []string {
	var result []string
	collector := func(s string) {
		result = append(result, s)
	}
	if tree.hasEmpty {
		result = append(result, "")
	}
	traverseVisitor(tree.head, collector)
	return result
}
