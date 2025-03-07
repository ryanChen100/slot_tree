package main

import (
	"fmt"
	"slot_tree/tree"
	"time"
)

var (
	// lineSetting = [][]int{{1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {2, 2, 2, 2, 2}, {0, 1, 2, 1, 0}, {2, 1, 0, 1, 2}}
	lineSetting = map[int][]int{0: {1, 1, 1, 1, 1}, 1: {0, 0, 0, 0, 0}, 2: {2, 2, 2, 2, 2}, 3: {0, 1, 2, 1, 0}, 4: {2, 1, 0, 1, 2}, 5: {1, 1, 1, 1, 2}, 6: {1, 1, 1, 1, 0},
		7: {0, 0, 0, 1, 0}, 8: {0, 0, 0, 1, 1}, 9: {0, 0, 0, 1, 2},
		10: {0, 0, 0, 2, 0}, 11: {0, 0, 0, 2, 1}, 12: {0, 0, 0, 2, 0},
		13: {0, 0, 0, 0, 1}, 14: {0, 0, 0, 0, 2}}
	reel  = [][]string{{"A", "A", "A"}, {"A", "A", "A"}, {"A", "B", "B"}, {"B", "B", "B"}, {"C", "C", "C"}}
	reel1 = [][]string{{"B", "B", "B"}, {"B", "B", "B"}, {"W", "W", "W"}, {"W", "W", "W"}, {"W", "W", "W"}}
	reel2 = [][]string{{"A", "B", "C"}, {"A", "W", "C"}, {"A", "B", "W"}, {"W", "B", "C"}, {"A", "B", "C"}}
	reel3 = [][]string{{"A", "A", "A"}, {"B", "B", "B"}, {"B", "B", "B"}, {"C", "C", "C"}, {"C", "C", "C"}}

	symbol = []string{}
	count  = []int{}

	slot = &tree.SlotData{
		PayLineMap: lineSetting,
	}
	slotMap = &tree.SlotData{
		PayLineMap: lineSetting,
	}
)

// ! pay line 補齊
func main() {
	slot.TreeInit()
	slot.Node.PrintTreeHorizontal("", true)

	time1 := time.Now()
	slot.Node.ReplaceReel(reel)
	fmt.Println("===============================")
	slot.PayData(slot.Node.TraverseLengthTree())
	fmt.Println("-----------reel end")
	slot.Node.ReplaceReel(reel1)
	slot.PayData(slot.Node.TraverseLengthTree())
	fmt.Println("-----------reel1 end")
	slot.Node.ReplaceReel(reel2)
	slot.PayData(slot.Node.TraverseLengthTree())
	fmt.Println("-----------reel2 end")
	slot.Node.ReplaceReel(reel3)
	slot.PayData(slot.Node.TraverseLengthTree())
	fmt.Println("-----------reel3 end")
	fmt.Printf("TraverseLengthTree : %v  \n", time.Since(time1))
}
