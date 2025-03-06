package tree

import "fmt"

// PayData 計算所有線得分
func (s *SlotData) PayData(treeInfo []ResultTree) {
	total := 0
	for _, tree := range treeInfo {
		tmp := s.payLine(tree)
		total += tmp
		fmt.Printf("Line: %d, Symbol: %v, Count: %d, Pay: %d\n",
			tree.lineIndex, tree.symbol, tree.count, tmp)
	}
	fmt.Println("total:", total)
}

// pay 計算單線得分
func (s *SlotData) payLine(tree ResultTree) int {
	//var sum int
	var tmpPay int
	var minSame bool
	if tree.symbol[0] != Wild {
		minSame = true
		goto checkPay
	}

	for i := 1; i < min; i++ {
		if tree.symbol[i] != tree.symbol[i-1] {
			minSame = true
			break
		}
	}

checkPay:
	if minSame {
		for i := 0; i < len(tree.symbol); i++ {
			if tree.symbol[i] == Wild {
				continue
			}
			for j := 0; j < len(s.SymbolList); j++ {
				if tree.symbol[i] == s.SymbolList[j] {
					tmpPay = s.PaySetting[tree.count-1][j]
				}
			}
		}
	} else {
		//最小長度都是W
		var wPay int
		for i := min; i < len(tree.symbol); i++ {
			//一樣是wild
			if tree.symbol[i] == Wild && len(tree.symbol) != i+1 {
				continue
			}

			for j := 0; j < len(s.SymbolList); j++ {
				if s.SymbolList[j] == Wild {
					wPay = s.PaySetting[i-1][j]
				}
			}

			for j := 0; j < len(s.SymbolList); j++ {
				if tree.symbol[i] == s.SymbolList[j] && tree.symbol[i] != Wild {
					tmpPay = s.PaySetting[tree.count-1][j]
				}
			}

			if wPay > tmpPay {
				tmpPay = wPay
			}

		}
	}

	return tmpPay
}
