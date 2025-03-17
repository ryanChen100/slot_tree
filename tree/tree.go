package tree

import "fmt"

// SlotData slot樹資料
type SlotData struct {
	Node           *node
	SymbolList     []string
	PayLineMap     map[int][]int
	payLineSetting [][]int //
	PaySetting     [][]int
	siteMap        map[int][]*node // 用來加速查找 site 節點
}

// Node 樹節點
type node struct {
	symbol    string  //結點圖案
	site      int     //結點位置
	children  []*node //子節點
	lineIndex []int   //pay line index

}

// ResultTree 擊中 樹節點結果
type ResultTree struct {
	count     int      // 中獎數量
	symbol    []string // 中獎符號
	lineIndex []int    // 中獎的支付線索引
}

var (
	//pay line最少擊中數
	min = 3
	// reel row
	row = 5

	//Wild 萬用symbol
	Wild = "W"
	//AllWild pay line 全部都是萬用symbol
	AllWild = "ALL Wild"
)

// TreeInit 樹初始化
func (s *SlotData) TreeInit() {
	newLine := make([][]int, len(s.PayLineMap))
	for index, lineSet := range s.PayLineMap {
		tmp := make([]int, len(lineSet))
		for i := 0; i < len(lineSet); i++ {
			tmp[i] = i*10 + lineSet[i]
		}
		newLine[index] = tmp
	}

	s.payLineSetting = newLine
	s.siteMap = make(map[int][]*node)
	s.Node = &node{site: -1}
	for lineIndex, payLineSet := range s.payLineSetting {
		s.Node.createTree(payLineSet, lineIndex)
	}
}

// TreeInitMap 樹初始化 使用site map 加速輪帶替換
func (s *SlotData) TreeInitMap() {
	newLine := make([][]int, len(s.PayLineMap))
	for index, lineSet := range s.PayLineMap {
		tmp := make([]int, len(lineSet))
		for i := 0; i < len(lineSet); i++ {
			tmp[i] = i*10 + lineSet[i]
		}
		newLine[index] = tmp
	}

	s.payLineSetting = newLine
	s.siteMap = make(map[int][]*node)
	s.Node = &node{site: -1}
	for lineIndex, payLineSet := range s.payLineSetting {
		s.Node.createTreeMap(payLineSet, lineIndex, s.siteMap)
	}
}

// createTreeMap 創建樹時代入siteMap
// 用遞歸的方式建立樹
func (n *node) createTreeMap(payLine []int, lineIndex int, siteMap map[int][]*node) *node {
	if len(payLine) == 0 {
		return nil
	}

	// 遍歷已有子節點，檢查是否已存在
	//共用節點
	for _, child := range n.children {
		if child.site == payLine[0] {
			// 確保子節點也記錄 lineIndex
			child.lineIndex = append(child.lineIndex, lineIndex)
			child.createTreeMap(payLine[1:], lineIndex, siteMap)
			return n
		}
	}

	// 若找不到對應節點，則創建新的節點
	//新建節點
	newNode := &node{
		site:      payLine[0],
		lineIndex: []int{lineIndex}, // 記錄支付線索引
	}
	n.children = append(n.children, newNode)
	// 存入 SiteMap
	siteMap[newNode.site] = append(siteMap[newNode.site], newNode)
	newNode.createTreeMap(payLine[1:], lineIndex, siteMap)

	return n
}

// createMap 創建樹
func (n *node) createTree(payLine []int, lineIndex int) *node {
	if len(payLine) == 0 {
		return nil
	}

	// 遍歷已有子節點，檢查是否已存在
	//共用節點
	for _, child := range n.children {
		if child.site == payLine[0] {
			// 確保子節點也記錄 lineIndex
			child.lineIndex = append(child.lineIndex, lineIndex)
			child.createTree(payLine[1:], lineIndex)
			return n
		}
	}

	// 若找不到對應節點，則創建新的節點
	//新建節點
	newNode := &node{
		site:      payLine[0],
		lineIndex: []int{lineIndex}, // 記錄支付線索引
	}
	n.children = append(n.children, newNode)
	newNode.createTree(payLine[1:], lineIndex)

	return n
}

// TraverseLengthTree 遍歷固定長度的樹並尋找目標值
func (n *node) TraverseLengthTree() []ResultTree {
	if n == nil {
		return nil
	}

	path := make([]string, 0) // 初始化 path
	return n.traverse(path)   // 傳遞 path 指標
}

func (n *node) traverse(path []string) []ResultTree {
	// 避免 root(-1) 被加入 path
	if n.site >= 0 {
		path = append(path, n.symbol) // 使用值傳遞，確保不影響其他遞迴分支
	}

	// 預計擊中數
	realLen := len(path) - 1

	// 優化 checkWild，減少多次函數呼叫
	if realLen > 0 {
		wildCheck1 := checkWild(path, len(path))
		wildCheck2 := checkWild(path, len(path)-1)

		if wildCheck1 != wildCheck2 && wildCheck2 != AllWild {
			if realLen >= min {
				return []ResultTree{{count: realLen, symbol: path[:realLen], lineIndex: n.lineIndex}}
			}
			return nil
		}
	}

	// 如果已經達到最大 row 數，回傳結果
	if len(path) == row {
		return []ResultTree{{count: len(path), symbol: path, lineIndex: n.lineIndex}}
	}

	// 遞迴遍歷子節點，預分配 res 容量
	res := make([]ResultTree, 0, len(n.children))
	for _, child := range n.children {
		res = append(res, child.traverse(path)...) // 傳遞新 slice 副本
	}

	return res
}

// ReplaceReel 替換輪帶 透過遞歸的方式一個一個查詢在替換
func (n *node) ReplaceReel(reel [][]string) {
	for i := range reel {
		siteX := i * 10
		for j := range reel[i] {
			n.replaceNode(siteX+j, reel[i][j])
		}
	}
}

// ReplaceReel 依賴於SlotData 替換輪帶 使用map方式替換欲取代位置
func (s *SlotData) ReplaceReel(reel [][]string) {
	for i := range reel {
		siteX := i * 10
		for j := range reel[i] {
			if node, inMap := s.siteMap[siteX+j]; inMap {
				for _, n := range node {
					n.symbol = reel[i][j]
				}
			}
		}
	}
}

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
