package tree

// replaceNode 替換節點值
func (n *node) replaceNode(site int, newValue string) {
	if n.site == site {
		n.symbol = newValue
	}

	for _, child := range n.children {
		child.replaceNode(site, newValue)
	}
}

func checkWild(line []string, index int) string {
	if index == 0 {
		return AllWild
	}
	if line[index-1] != Wild {
		return line[index-1]
	}

	return checkWild(line, index-1)
}
