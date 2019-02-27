package ternarytree

// TernaryTree is a ternary search tree data structure.
type TernaryTree struct {
	head     *treeNode
	terminal bool
}

type treeNode struct {
	char     byte
	loKid    *treeNode
	eqKid    *treeNode
	hiKid    *treeNode
	terminal bool
}

// Insert adds a string to the ternary search tree.
func (tree *TernaryTree) Insert(s string) {
	b := []byte(s)
	if len(b) > 0 {
		tree.head = insertVisitor(tree.head, b[0], b[1:])
	} else {
		tree.terminal = true
	}
}

// Search tests whether a string was added to the ternary search tree.
func (tree *TernaryTree) Search(s string) bool {
	b := []byte(s)
	if len(b) > 0 {
		return searchVisitor(tree.head, b[0], b[1:])
	} else if tree.terminal {
		return true
	} else {
		return false
	}
}

func insertVisitor(node *treeNode, head byte, tail []byte) *treeNode {
	if node == nil {
		node = &treeNode{char: head}
	}
	if head < node.char {
		node.loKid = insertVisitor(node.loKid, head, tail)
	} else if head == node.char {
		if len(tail) > 0 {
			node.eqKid = insertVisitor(node.eqKid, tail[0], tail[1:])
		} else {
			node.terminal = true
		}
	} else {
		node.hiKid = insertVisitor(node.hiKid, head, tail)
	}
	return node
}

func searchVisitor(node *treeNode, head byte, tail []byte) bool {
	if node == nil {
		return false
	}
	if head < node.char {
		return searchVisitor(node.loKid, head, tail)
	} else if head == node.char {
		if len(tail) > 0 {
			return searchVisitor(node.eqKid, tail[0], tail[1:])
		}
		return node.terminal
	} else {
		return searchVisitor(node.hiKid, head, tail)
	}
}
