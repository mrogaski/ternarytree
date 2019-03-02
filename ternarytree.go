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

// Search tests whether a string was added to the ternary search tree.
func (tree *TernaryTree) Search(s string) bool {
	if len(s) > 0 {
		return searchVisitor(tree.head, s[0], s[1:])
	} else if tree.hasEmpty {
		return true
	} else {
		return false
	}
}

func searchVisitor(node *treeNode, head byte, tail string) bool {
	if node == nil {
		return false
	}
	if head < node.char {
		return searchVisitor(node.loKid, head, tail)
	} else if head == node.char {
		if len(tail) > 0 {
			return searchVisitor(node.eqKid, tail[0], tail[1:])
		}
		return node.data != nil
	} else {
		return searchVisitor(node.hiKid, head, tail)
	}
}

// PartialMatchSearch finds matches for wildcard patterns in the ternary search tree.
func (tree *TernaryTree) PartialMatchSearch(s string, wildcard byte) []string {
	var result []string
	if len(s) > 0 {
		pmSearchVisitor(tree.head, s[0], s[1:], &result)
	} else if tree.hasEmpty {
		result = append(result, "")
	}
	return result
}

func pmSearchVisitor(node *treeNode, head byte, tail string, result *[]string) {}
