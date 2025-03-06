package tree

import "testing"

// go test -bench=. -benchtime=1000000x
// replaceReel Map方式優於loop

// 生成測試用的 SlotData
func setupTestSlotData() *SlotData {
	s := &SlotData{
		PayLineMap: map[int][]int{
			0: {0, 10, 20, 30, 40},
			1: {1, 11, 21, 31, 41},
			2: {2, 12, 22, 32, 42},
		},
	}

	return s
}

// 生成測試用的 Reel
func setupTestReel() [][]string {
	return [][]string{
		{"A", "B", "C", "D", "E"},
		{"F", "G", "H", "I", "J"},
		{"K", "L", "M", "N", "O"},
	}
}

// Benchmark - 測試使用 `siteMap` 直接修改 symbol 的效能
func BenchmarkReplaceReelWithMap(b *testing.B) {
	s := setupTestSlotData()
	s.TreeInitMap()
	reel := setupTestReel()

	b.ResetTimer() // 重置計時器
	for i := 0; i < b.N; i++ {
		s.ReplaceReel(reel)
	}
}

// Benchmark - 測試遞迴方式修改 symbol 的效能
func BenchmarkReplaceReelRecursive(b *testing.B) {
	s := setupTestSlotData()
	s.TreeInit()
	reel := setupTestReel()

	b.ResetTimer() // 重置計時器
	for i := 0; i < b.N; i++ {
		s.Node.ReplaceReel(reel)
	}
}
