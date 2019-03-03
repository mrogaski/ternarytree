package ternarytree

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
		pmSearchVisitor(tree.head, wildcard, s[0], s[1:], &result)
	} else if tree.hasEmpty {
		result = append(result, "")
	}
	return result
}

func pmSearchVisitor(node *treeNode, wildcard byte, head byte, tail string, result *[]string) {
	if node == nil {
		return
	}
	if head == wildcard || head < node.char {
		pmSearchVisitor(node.loKid, wildcard, head, tail, result)
	}
	if head == wildcard || head == node.char {
		if len(tail) > 0 {
			pmSearchVisitor(node.eqKid, wildcard, tail[0], tail[1:], result)
		} else {
			*result = append(*result, *node.data)
		}
	}
	if head == wildcard || head > node.char {
		pmSearchVisitor(node.hiKid, wildcard, head, tail, result)
	}
}

// NearNeighborSearch finds strings withing a given Hamming distance of the target string.
func (tree *TernaryTree) NearNeighborSearch(s string, distance int) []string {
	var result []string
	if len(s) > 0 {
		nnSearchVisitor(tree.head, distance, s[0], s[1:], &result)
	} else if tree.hasEmpty {
		result = append(result, "")
	}
	return result
}

func nnSearchVisitor(node *treeNode, distance int, head byte, tail string, result *[]string) {
	if node == nil || distance < 0 {
		return
	}

	if distance > 0 || head < node.char {
		nnSearchVisitor(node.loKid, distance, head, tail, result)
	}

	d := distance
	if head != node.char {
		d = distance - 1
	}
	if node.data != nil {
		if len(tail) <= d {
			*result = append(*result, *node.data)
		}
	} else {
		if len(tail) > 0 {
			nnSearchVisitor(node.eqKid, d, tail[0], tail[1:], result)
		}
	}

	if distance > 0 || head > node.char {
		nnSearchVisitor(node.hiKid, distance, head, tail, result)
	}
}
