package ternarytree

// Insert adds a string to the ternary search tree.
func (tree *TernaryTree) Insert(s string) {
	if len(s) > 0 {
		tree.head = insertVisitor(tree.head, s[0], s[1:], &s)
	} else {
		tree.hasEmpty = true
	}
}

func insertVisitor(node *treeNode, head byte, tail string, start *string) *treeNode {
	if node == nil {
		node = &treeNode{char: head}
	}
	if head < node.char {
		node.loKid = insertVisitor(node.loKid, head, tail, start)
	} else if head == node.char {
		if len(tail) > 0 {
			node.eqKid = insertVisitor(node.eqKid, tail[0], tail[1:], start)
		} else {
			node.data = start
		}
	} else {
		node.hiKid = insertVisitor(node.hiKid, head, tail, start)
	}
	return node
}
