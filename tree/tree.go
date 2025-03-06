package tree

// SlotData slot樹資料
type SlotData struct {
	Node           *node
	SymbolList     []string
	PayLineMap     map[int][]int
	payLineSetting [][]int //
	PaySetting     [][]int
	siteMap map[int][]*node // 用來加速查找 site 節點
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

	// 初始化 path 和 sitePath
	var path []string
	var sitePath []int

	return n.traverse(path, sitePath)
}

// traverse 遞迴遍歷樹
func (n *node) traverse(path []string, sitePath []int) []ResultTree {
	// 去掉根節點
	if n.site >= 0 {
		path = append(path, n.symbol)
		sitePath = append(sitePath, n.site)
	}

	// 至少中 min，進入多一輪才能判斷實際擊中數
	realLen := len(path) - 1
	if len(path) > 1 {
		if checkWild(path, len(path)) != checkWild(path, len(path)-1) && checkWild(path, len(path)-1) != AllWild {
			if realLen >= min {
				return []ResultTree{{count: realLen, symbol: path[:realLen], lineIndex: n.lineIndex}}
			}
			return nil
		}
	}

	// 全中
	if len(path) == row {
		return []ResultTree{{count: len(path), symbol: path, lineIndex: n.lineIndex}}
	}

	var res []ResultTree
	for _, child := range n.children {
		res = append(res, child.traverse(path, sitePath)...)
	}
	return res
}

// TraverseLengthTreeFast 提高算線方式
// TODO 待驗證
// func (s *SlotData) TraverseLengthTreeFast() []ResultTree {
// 	if s.Node == nil {
// 		return nil
// 	}

// 	var results []ResultTree

// 	// 遍歷所有支付線
// 	for lineIndex, lineSet := range s.payLineSetting {
// 		var path []string
// 		var sitePath []int

// 		// 直接按照支付線順序抓取符號
// 		for _, pos := range lineSet {
// 			node := s.findNodeBySite(s.Node, pos)
// 			if node != nil {
// 				path = append(path, node.symbol)
// 				sitePath = append(sitePath, node.site)
// 			}
// 		}

// 		// 如果路徑長度小於 min，直接跳過
// 		if len(path) < min {
// 			continue
// 		}

// 		// 檢查中獎條件
// 		count := checkWinningPattern(path)
// 		if count >= min {
// 			// 使用 cache 加速
// 			key := fmt.Sprintf("%v", path)
// 			if val, found := s.Cache[key]; found {
// 				results = append(results, val)
// 			} else {
// 				result := ResultTree{count: count, symbol: path[:count], lineIndex: lineIndex}
// 				s.Cache[key] = result
// 				results = append(results, result)
// 			}
// 		}
// 	}
// 	return results
// }

func (s *SlotData) findNodeBySite(n *node, site int) *node {
	if n == nil {
		return nil
	}
	if n.site == site {
		return n
	}
	for _, child := range n.children {
		if res := s.findNodeBySite(child, site); res != nil {
			return res
		}
	}
	return nil
}

func checkWinningPattern(path []string) int {
	if len(path) < min {
		return 0
	}

	count := 1
	for i := 1; i < len(path); i++ {
		if path[i] == path[i-1] || path[i] == Wild || path[i-1] == Wild {
			count++
		} else {
			break
		}
	}

	return count
}

// ReplaceReel 替換輪帶
func (n *node) ReplaceReel(reel [][]string) {
	for i := range reel {
		siteX := i * 10
		for j := range reel[i] {
			n.replaceNode(siteX+j, reel[i][j])
		}
	}
}

// ReplaceReel 替換輪帶
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
