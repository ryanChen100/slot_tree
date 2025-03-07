package tree

import "fmt"

// replaceNode 替換節點值
func (n *node) replaceNode(site int, newValue string) {
	if n.site == site {
		n.symbol = newValue
	}

	for _, child := range n.children {
		child.replaceNode(site, newValue)
	}
}

// func checkWild(line []string, index int) string {
// 	if index == 0 {
// 		return AllWild
// 	}
// 	if line[index-1] != Wild {
// 		return line[index-1]
// 	}

// 	return checkWild(line, index-1)
// }

func checkWild(line []string, index int) string {
	if index == 0 {
		return AllWild
	}

	// **使用迴圈替代遞迴**
	for i := index - 1; i >= 0; i-- {
		if line[i] != Wild {
			return line[i] // 找到第一個非 Wild 的符號
		}
	}

	return AllWild // 整條線都是 Wild
}

// 水平打印樹
func (n *node) PrintTreeHorizontal(prefix string, isLast bool) {
	if n == nil {
		return
	}

	// 定義樹的縮排格式
	indent := prefix
	if isLast {
		indent += "   " // 最後一個子節點不需要 `│`
	} else {
		indent += "│  " // 其他子節點需要 `│` 來對齊
	}

	// 印出當前節點，包含 `lineIndex`
	fmt.Printf("%s├── Site: %d, LineIndex: %d\n", prefix, n.site, n.lineIndex)

	// 遞迴遍歷子節點
	for i, child := range n.children {
		child.PrintTreeHorizontal(indent, i == len(n.children)-1)
	}
}
